package blackhole_test

import (
	"context"
	"testing"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/buf"
	"github.com/xmplusdev/xmcore/common/serial"
	"github.com/xmplusdev/xmcore/common/session"
	"github.com/xmplusdev/xmcore/proxy/blackhole"
	"github.com/xmplusdev/xmcore/transport"
	"github.com/xmplusdev/xmcore/transport/pipe"
)

func TestBlackholeHTTPResponse(t *testing.T) {
	ctx := session.ContextWithOutbounds(context.Background(), []*session.Outbound{{}})
	handler, err := blackhole.New(ctx, &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	var mb buf.MultiBuffer
	var rerr error
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		rerr = e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(ctx, &link, nil))
	common.Must(rerr)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
