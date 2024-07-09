package httpupgrade

import (
	"bufio"
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/common/session"
	"github.com/xmplusdev/xmcore/transport/internet"
	"github.com/xmplusdev/xmcore/transport/internet/stat"
	"github.com/xmplusdev/xmcore/transport/internet/tls"
)

type ConnRF struct {
	net.Conn
	Req   *http.Request
	First bool
}

func (c *ConnRF) Read(b []byte) (int, error) {
	if c.First {
		c.First = false
		// create reader capped to size of `b`, so it can be fully drained into
		// `b` later with a single Read call
		reader := bufio.NewReaderSize(c.Conn, len(b))
		resp, err := http.ReadResponse(reader, c.Req) // nolint:bodyclose
		if err != nil {
			return 0, err
		}
		if resp.Status != "101 Switching Protocols" ||
			strings.ToLower(resp.Header.Get("Upgrade")) != "websocket" ||
			strings.ToLower(resp.Header.Get("Connection")) != "upgrade" {
			return 0, newError("unrecognized reply")
		}
		// drain remaining bufreader
		return reader.Read(b[:reader.Buffered()])
	}
	return c.Conn.Read(b)
}

func dialhttpUpgrade(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (net.Conn, error) {
	transportConfiguration := streamSettings.ProtocolSettings.(*Config)

	pconn, err := internet.DialSystem(ctx, dest, streamSettings.SocketSettings)
	if err != nil {
		newError("failed to dial to ", dest).Base(err).AtError().WriteToLog()
		return nil, err
	}

	var conn net.Conn
	var requestURL url.URL
	if config := tls.ConfigFromStreamSettings(streamSettings); config != nil {
		tlsConfig := config.GetTLSConfig(tls.WithDestination(dest), tls.WithNextProto("http/1.1"))
		if fingerprint := tls.GetFingerprint(config.Fingerprint); fingerprint != nil {
			conn = tls.UClient(pconn, tlsConfig, fingerprint)
			if err := conn.(*tls.UConn).WebsocketHandshakeContext(ctx); err != nil {
				return nil, err
			}
		} else {
			conn = tls.Client(pconn, tlsConfig)
		}
		requestURL.Scheme = "https"
	} else {
		conn = pconn
		requestURL.Scheme = "http"
	}

	requestURL.Host = dest.NetAddr()
	requestURL.Path = transportConfiguration.GetNormalizedPath()
	req := &http.Request{
		Method: http.MethodGet,
		URL:    &requestURL,
		Host:   transportConfiguration.Host,
		Header: make(http.Header),
	}
	for key, value := range transportConfiguration.Header {
		AddHeader(req.Header, key, value)
	}
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")

	err = req.Write(conn)
	if err != nil {
		return nil, err
	}

	connRF := &ConnRF{
		Conn:  conn,
		Req:   req,
		First: true,
	}

	if transportConfiguration.Ed == 0 {
		_, err = connRF.Read([]byte{})
		if err != nil {
			return nil, err
		}
	}

	return connRF, nil
}

//http.Header.Add() will convert headers to MIME header format.
//Some people don't like this because they want to send "Web*S*ocket".
//So we add a simple function to replace that method.
func AddHeader(header http.Header, key, value string) {
	header[key] = append(header[key], value)
}

func Dial(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
	newError("creating connection to ", dest).WriteToLog(session.ExportIDToError(ctx))

	conn, err := dialhttpUpgrade(ctx, dest, streamSettings)
	if err != nil {
		return nil, newError("failed to dial request to ", dest).Base(err)
	}
	return stat.Connection(conn), nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, Dial))
}
