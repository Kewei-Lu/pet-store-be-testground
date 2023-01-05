package cmd

import (
	"context"

	"github.com/gogf/csrf/v2"
	"net/http"
	"petStore/internal/controller"
	"time"
	// "github.com/goflyfox/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// s.BindHandler("/cookie", func(r *ghttp.Request) {
			// 	datetime := r.Cookie.Get("datetime")
			// 	r.Cookie.Set("datetime", gtime.Datetime())
			// 	r.Response.Write("datetime:", datetime)
			// })
			s.Group("/api", func(group *ghttp.RouterGroup) {

				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					// gtoken.Middleware(group)
					group.Bind(
						controller.User,
					)
				})
				group.Group("/money", func(group *ghttp.RouterGroup) {
					group.Middleware(csrf.NewWithCfg(csrf.Config{
						Cookie: &http.Cookie{
							Name: "X-Token", // token name in cookie
						},
						ExpireTime:      time.Hour * 24,
						TokenLength:     32,
						TokenRequestKey: "X-Token", // use this key to read token in request param
					}))
					group.Bind(
						controller.Money,
					)
				})
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Money,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
