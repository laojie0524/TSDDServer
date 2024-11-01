package base

import (
	"embed"

	"github.com/TangSengDaoDao/TangSengDaoDaoServer/modules/base/app"
	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/register"
)

//go:embed sql
var sqlFS embed.FS

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			SetupAPI: func() register.APIRouter {
				return app.New(ctx.(*config.Context))
			},
			SQLDir: register.NewSQLFS(sqlFS),
		}
	})
}
