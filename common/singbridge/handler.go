package singbridge

import (
	"context"
	"io"

	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"github.com/xmplusdev/xmcore/common/buf"
	"github.com/xmplusdev/xmcore/common/errors"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/features/routing"
	"github.com/xmplusdev/xmcore/transport"
)

var (
	_ N.TCPConnectionHandler = (*Dispatcher)(nil)
	_ N.UDPConnectionHandler = (*Dispatcher)(nil)
)

type Dispatcher struct {
	upstream     routing.Dispatcher
	newErrorFunc func(values ...any) *errors.Error
}

func NewDispatcher(dispatcher routing.Dispatcher, newErrorFunc func(values ...any) *errors.Error) *Dispatcher {
	return &Dispatcher{
		upstream:     dispatcher,
		newErrorFunc: newErrorFunc,
	}
}

func (d *Dispatcher) NewConnection(ctx context.Context, conn net.Conn, metadata M.Metadata) error {
	xConn := NewConn(conn)
	return d.upstream.DispatchLink(ctx, ToDestination(metadata.Destination, net.Network_TCP), &transport.Link{
		Reader: xConn,
		Writer: xConn,
	})
}

func (d *Dispatcher) NewPacketConnection(ctx context.Context, conn N.PacketConn, metadata M.Metadata) error {
	return d.upstream.DispatchLink(ctx, ToDestination(metadata.Destination, net.Network_UDP), &transport.Link{
		Reader: buf.NewPacketReader(conn.(io.Reader)),
		Writer: buf.NewWriter(conn.(io.Writer)),
	})
}

func (d *Dispatcher) NewError(ctx context.Context, err error) {
	errors.LogInfo(ctx, err.Error())
}
