package statistics

import (
	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/register"
)

func init() {
	register.AddModule(func(ctx interface{}) register.Module {
		x := ctx.(*config.Context)
		return register.Module{
			Name: "statistics",
			SetupAPI: func() register.APIRouter {
				return NewStatistics(x)
			},
		}
	})
}
