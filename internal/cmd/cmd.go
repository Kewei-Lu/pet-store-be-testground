package cmd

import (
	"context"

	"net/http"
	"petStore/internal/controller"
	"time"

	"github.com/gogf/csrf/v2"

	// "github.com/goflyfox/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var JwtConf = csrf.NewWithCfg(csrf.Config{
	Cookie: &http.Cookie{
		Name: "jwt-token", // token name in cookie
	},
	ExpireTime:      time.Hour * 24,
	TokenLength:     32,
	TokenRequestKey: "jwt-token", // use this key to read token in request param
})

var XTokenConf = csrf.NewWithCfg(csrf.Config{
	Cookie: &http.Cookie{
		Name: "X-Token", // token name in cookie
	},
	ExpireTime:      time.Hour * 24,
	TokenLength:     32,
	TokenRequestKey: "X-Token", // use this key to read token in request param
})

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					// add XTOKEN to prevent csrf
					group.Middleware(XTokenConf)
					group.Bind(
						controller.User,
					)
				})
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(JwtConf)
					group.Group("/money", func(group *ghttp.RouterGroup) {
						group.Bind(
							controller.Money,
						)
					})
					group.Group("/auth", func(group *ghttp.RouterGroup) {
						group.Bind(
							controller.Auth,
						)
					})
				})

			})
			s.Run()
			return nil
		},
	}
)
