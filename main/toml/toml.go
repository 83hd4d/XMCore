package toml

import (
	"context"
	"io"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/cmdarg"
	"github.com/xmplusdev/xmcore/common/errors"
	"github.com/xmplusdev/xmcore/core"
	"github.com/xmplusdev/xmcore/infra/conf"
	"github.com/xmplusdev/xmcore/infra/conf/serial"
	"github.com/xmplusdev/xmcore/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "TOML",
		Extension: []string{"toml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					errors.LogInfo(context.Background(), "Reading config: ", arg)
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, errors.New("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeTOMLConfig(r)
					if err != nil {
						return nil, errors.New("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadTOMLConfig(v)
			default:
				return nil, errors.New("unknown type")
			}
		},
	}))
}
