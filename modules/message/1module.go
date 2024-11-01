package message

import (
	"embed"

	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/register"
)

//go:embed sql
var sqlFS embed.FS

//go:embed swagger/api.yaml
var swaggerContent string

//go:embed swagger/conversation.yaml
var conversationSwagger string

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			Name: "message",
			SetupAPI: func() register.APIRouter {
				return New(ctx.(*config.Context))
			},
			SQLDir:  register.NewSQLFS(sqlFS),
			Swagger: swaggerContent,
		}
	})

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			Name: "conversation",
			SetupAPI: func() register.APIRouter {
				return NewConversation(ctx.(*config.Context))
			},
			Swagger: conversationSwagger,
		}
	})
	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			SetupAPI: func() register.APIRouter {
				return NewManager(ctx.(*config.Context))
			},
		}
	})
}
