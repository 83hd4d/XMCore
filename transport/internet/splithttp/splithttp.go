package splithttp

import (
	"context"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/errors"
)

//go:generate go run github.com/xmplusdev/xmcore/common/errors/errorgen

const protocolName = "splithttp"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, errors.New("splithttp is a transport protocol.")
	}))
}
