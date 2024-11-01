package robot

import (
	"embed"

	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/register"
)

//go:embed sql
var sqlFS embed.FS

//go:embed swagger/api.yaml
var swaggerContent string

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			Name: "robot",
			SetupAPI: func() register.APIRouter {
				return New(ctx.(*config.Context))
			},
			SQLDir:  register.NewSQLFS(sqlFS),
			Swagger: swaggerContent,
		}
	})
}