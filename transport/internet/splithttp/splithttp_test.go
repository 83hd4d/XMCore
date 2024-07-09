package splithttp_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/common/protocol/tls/cert"
	"github.com/xmplusdev/xmcore/testing/servers/tcp"
	"github.com/xmplusdev/xmcore/transport/internet"
	. "github.com/xmplusdev/xmcore/transport/internet/splithttp"
	"github.com/xmplusdev/xmcore/transport/internet/stat"
	"github.com/xmplusdev/xmcore/transport/internet/tls"
)

func Test_listenSHAndDial(t *testing.T) {
	listenPort := tcp.PickPort()
	listen, err := ListenSH(context.Background(), net.LocalHostIP, listenPort, &internet.MemoryStreamConfig{
		ProtocolName: "splithttp",
		ProtocolSettings: &Config{
			Path: "/sh",
		},
	}, func(conn stat.Connection) {
		go func(c stat.Connection) {
			defer c.Close()

			var b [1024]byte
			_, err := c.Read(b[:])
			if err != nil {
				return
			}

			common.Must2(c.Write([]byte("Response")))
		}(conn)
	})
	common.Must(err)
	ctx := context.Background()
	streamSettings := &internet.MemoryStreamConfig{
		ProtocolName:     "splithttp",
		ProtocolSettings: &Config{Path: "sh"},
	}
	conn, err := Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), listenPort), streamSettings)

	common.Must(err)
	_, err = conn.Write([]byte("Test connection 1"))
	common.Must(err)

	var b [1024]byte
	fmt.Println("test2")
	n, _ := conn.Read(b[:])
	fmt.Println("string is", n)
	if string(b[:n]) != "Response" {
		t.Error("response: ", string(b[:n]))
	}

	common.Must(conn.Close())
	<-time.After(time.Second * 5)
	conn, err = Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), listenPort), streamSettings)
	common.Must(err)
	_, err = conn.Write([]byte("Test connection 2"))
	common.Must(err)
	n, _ = conn.Read(b[:])
	common.Must(err)
	if string(b[:n]) != "Response" {
		t.Error("response: ", string(b[:n]))
	}
	common.Must(conn.Close())

	common.Must(listen.Close())
}

func TestDialWithRemoteAddr(t *testing.T) {
	listenPort := tcp.PickPort()
	listen, err := ListenSH(context.Background(), net.LocalHostIP, listenPort, &internet.MemoryStreamConfig{
		ProtocolName: "splithttp",
		ProtocolSettings: &Config{
			Path: "sh",
		},
	}, func(conn stat.Connection) {
		go func(c stat.Connection) {
			defer c.Close()

			var b [1024]byte
			_, err := c.Read(b[:])
			// common.Must(err)
			if err != nil {
				return
			}

			_, err = c.Write([]byte("Response"))
			common.Must(err)
		}(conn)
	})
	common.Must(err)

	conn, err := Dial(context.Background(), net.TCPDestination(net.DomainAddress("localhost"), listenPort), &internet.MemoryStreamConfig{
		ProtocolName:     "splithttp",
		ProtocolSettings: &Config{Path: "sh", Header: map[string]string{"X-Forwarded-For": "1.1.1.1"}},
	})

	common.Must(err)
	_, err = conn.Write([]byte("Test connection 1"))
	common.Must(err)

	var b [1024]byte
	n, _ := conn.Read(b[:])
	if string(b[:n]) != "Response" {
		t.Error("response: ", string(b[:n]))
	}

	common.Must(listen.Close())
}

func Test_listenSHAndDial_TLS(t *testing.T) {
	if runtime.GOARCH == "arm64" {
		return
	}

	listenPort := tcp.PickPort()

	start := time.Now()

	streamSettings := &internet.MemoryStreamConfig{
		ProtocolName: "splithttp",
		ProtocolSettings: &Config{
			Path: "shs",
		},
		SecurityType: "tls",
		SecuritySettings: &tls.Config{
			AllowInsecure: true,
			Certificate:   []*tls.Certificate{tls.ParseCertificate(cert.MustGenerate(nil, cert.CommonName("localhost")))},
		},
	}
	listen, err := ListenSH(context.Background(), net.LocalHostIP, listenPort, streamSettings, func(conn stat.Connection) {
		go func() {
			_ = conn.Close()
		}()
	})
	common.Must(err)
	defer listen.Close()

	conn, err := Dial(context.Background(), net.TCPDestination(net.DomainAddress("localhost"), listenPort), streamSettings)
	common.Must(err)
	_ = conn.Close()

	end := time.Now()
	if !end.Before(start.Add(time.Second * 5)) {
		t.Error("end: ", end, " start: ", start)
	}
}
