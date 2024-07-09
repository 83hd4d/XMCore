package loopback

import (
	"context"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/buf"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/common/net/cnc"
	"github.com/xmplusdev/xmcore/common/retry"
	"github.com/xmplusdev/xmcore/common/session"
	"github.com/xmplusdev/xmcore/common/task"
	"github.com/xmplusdev/xmcore/core"
	"github.com/xmplusdev/xmcore/features/routing"
	"github.com/xmplusdev/xmcore/transport"
	"github.com/xmplusdev/xmcore/transport/internet"
)

type Loopback struct {
	config             *Config
	dispatcherInstance routing.Dispatcher
}

func (l *Loopback) Process(ctx context.Context, link *transport.Link, _ internet.Dialer) error {
	outbounds := session.OutboundsFromContext(ctx)
	ob := outbounds[len(outbounds) - 1]
	if !ob.Target.IsValid() {
		return newError("target not specified.")
	}
	ob.Name = "loopback"
	destination := ob.Target

	newError("opening connection to ", destination).WriteToLog(session.ExportIDToError(ctx))

	input := link.Reader
	output := link.Writer

	var conn net.Conn
	err := retry.ExponentialBackoff(2, 100).On(func() error {
		dialDest := destination

		content := new(session.Content)
		content.SkipDNSResolve = true

		ctx = session.ContextWithContent(ctx, content)

		inbound := session.InboundFromContext(ctx)

		inbound.Tag = l.config.InboundTag

		ctx = session.ContextWithInbound(ctx, inbound)

		rawConn, err := l.dispatcherInstance.Dispatch(ctx, dialDest)
		if err != nil {
			return err
		}

		var readerOpt cnc.ConnectionOption
		if dialDest.Network == net.Network_TCP {
			readerOpt = cnc.ConnectionOutputMulti(rawConn.Reader)
		} else {
			readerOpt = cnc.ConnectionOutputMultiUDP(rawConn.Reader)
		}

		conn = cnc.NewConnection(cnc.ConnectionInputMulti(rawConn.Writer), readerOpt)
		return nil
	})
	if err != nil {
		return newError("failed to open connection to ", destination).Base(err)
	}
	defer conn.Close()

	requestDone := func() error {
		var writer buf.Writer
		if destination.Network == net.Network_TCP {
			writer = buf.NewWriter(conn)
		} else {
			writer = &buf.SequentialWriter{Writer: conn}
		}

		if err := buf.Copy(input, writer); err != nil {
			return newError("failed to process request").Base(err)
		}

		return nil
	}

	responseDone := func() error {
		var reader buf.Reader
		if destination.Network == net.Network_TCP {
			reader = buf.NewReader(conn)
		} else {
			reader = buf.NewPacketReader(conn)
		}
		if err := buf.Copy(reader, output); err != nil {
			return newError("failed to process response").Base(err)
		}

		return nil
	}

	if err := task.Run(ctx, requestDone, task.OnSuccess(responseDone, task.Close(output))); err != nil {
		return newError("connection ends").Base(err)
	}

	return nil
}

func (l *Loopback) init(config *Config, dispatcherInstance routing.Dispatcher) error {
	l.dispatcherInstance = dispatcherInstance
	l.config = config
	return nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		l := new(Loopback)
		err := core.RequireFeatures(ctx, func(dispatcherInstance routing.Dispatcher) error {
			return l.init(config.(*Config), dispatcherInstance)
		})
		return l, err
	}))
}
