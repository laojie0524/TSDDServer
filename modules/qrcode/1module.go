package qrcode

import (
	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/register"
)

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			SetupAPI: func() register.APIRouter {
				return New(ctx.(*config.Context))
			},
		}
	})
}
