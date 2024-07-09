package udp

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/buf"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/common/protocol/udp"
	"github.com/xmplusdev/xmcore/common/session"
	"github.com/xmplusdev/xmcore/common/signal"
	"github.com/xmplusdev/xmcore/common/signal/done"
	"github.com/xmplusdev/xmcore/features/routing"
	"github.com/xmplusdev/xmcore/transport"
)

type ResponseCallback func(ctx context.Context, packet *udp.Packet)

type connEntry struct {
	link   *transport.Link
	timer  signal.ActivityUpdater
	cancel context.CancelFunc
}

type Dispatcher struct {
	sync.RWMutex
	conn       *connEntry
	dispatcher routing.Dispatcher
	callback   ResponseCallback
	callClose  func() error
}

func NewDispatcher(dispatcher routing.Dispatcher, callback ResponseCallback) *Dispatcher {
	return &Dispatcher{
		dispatcher: dispatcher,
		callback:   callback,
	}
}

func (v *Dispatcher) RemoveRay() {
	v.Lock()
	defer v.Unlock()
	if v.conn != nil {
		common.Interrupt(v.conn.link.Reader)
		common.Close(v.conn.link.Writer)
		v.conn = nil
	}
}

func (v *Dispatcher) getInboundRay(ctx context.Context, dest net.Destination) (*connEntry, error) {
	v.Lock()
	defer v.Unlock()

	if v.conn != nil {
		return v.conn, nil
	}

	newError("establishing new connection for ", dest).WriteToLog()

	ctx, cancel := context.WithCancel(ctx)
	removeRay := func() {
		cancel()
		v.RemoveRay()
	}
	timer := signal.CancelAfterInactivity(ctx, removeRay, time.Minute)

	link, err := v.dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return nil, newError("failed to dispatch request to ", dest).Base(err)
	}

	entry := &connEntry{
		link:   link,
		timer:  timer,
		cancel: removeRay,
	}
	v.conn = entry
	go handleInput(ctx, entry, dest, v.callback, v.callClose)
	return entry, nil
}

func (v *Dispatcher) Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer) {
	// TODO: Add user to destString
	newError("dispatch request to: ", destination).AtDebug().WriteToLog(session.ExportIDToError(ctx))

	conn, err := v.getInboundRay(ctx, destination)
	if err != nil {
		newError("failed to get inbound").Base(err).WriteToLog(session.ExportIDToError(ctx))
		return
	}
	outputStream := conn.link.Writer
	if outputStream != nil {
		if err := outputStream.WriteMultiBuffer(buf.MultiBuffer{payload}); err != nil {
			newError("failed to write first UDP payload").Base(err).WriteToLog(session.ExportIDToError(ctx))
			conn.cancel()
			return
		}
	}
}

func handleInput(ctx context.Context, conn *connEntry, dest net.Destination, callback ResponseCallback, callClose func() error) {
	defer func() {
		conn.cancel()
		if callClose != nil {
			callClose()
		}
	}()

	input := conn.link.Reader
	timer := conn.timer

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		mb, err := input.ReadMultiBuffer()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				newError("failed to handle UDP input").Base(err).WriteToLog(session.ExportIDToError(ctx))
			}
			return
		}
		timer.Update()
		for _, b := range mb {
			if b.UDP != nil {
				dest = *b.UDP
			}
			callback(ctx, &udp.Packet{
				Payload: b,
				Source:  dest,
			})
		}
	}
}

type dispatcherConn struct {
	dispatcher *Dispatcher
	cache      chan *udp.Packet
	done       *done.Instance
	ctx        context.Context
}

func DialDispatcher(ctx context.Context, dispatcher routing.Dispatcher) (net.PacketConn, error) {
	c := &dispatcherConn{
		cache: make(chan *udp.Packet, 16),
		done:  done.New(),
		ctx:   ctx,
	}

	d := &Dispatcher{
		dispatcher: dispatcher,
		callback:   c.callback,
		callClose:  c.Close,
	}
	c.dispatcher = d
	return c, nil
}

func (c *dispatcherConn) callback(ctx context.Context, packet *udp.Packet) {
	select {
	case <-c.done.Wait():
		packet.Payload.Release()
		return
	case c.cache <- packet:
	default:
		packet.Payload.Release()
		return
	}
}

func (c *dispatcherConn) ReadFrom(p []byte) (int, net.Addr, error) {
	var packet *udp.Packet
s:
	select {
	case <-c.done.Wait():
		select {
		case packet = <-c.cache:
			break s
		default:
			return 0, nil, io.EOF
		}
	case packet = <-c.cache:
	}
	return copy(p, packet.Payload.Bytes()), &net.UDPAddr{
		IP:   packet.Source.Address.IP(),
		Port: int(packet.Source.Port),
	}, nil
}

func (c *dispatcherConn) WriteTo(p []byte, addr net.Addr) (int, error) {
	buffer := buf.New()
	raw := buffer.Extend(buf.Size)
	n := copy(raw, p)
	buffer.Resize(0, int32(n))

	destination := net.DestinationFromAddr(addr)
	buffer.UDP = &destination
	c.dispatcher.Dispatch(c.ctx, destination, buffer)
	return n, nil
}

func (c *dispatcherConn) Close() error {
	return c.done.Close()
}

func (c *dispatcherConn) LocalAddr() net.Addr {
	return &net.UDPAddr{
		IP:   []byte{0, 0, 0, 0},
		Port: 0,
	}
}

func (c *dispatcherConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *dispatcherConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *dispatcherConn) SetWriteDeadline(t time.Time) error {
	return nil
}
