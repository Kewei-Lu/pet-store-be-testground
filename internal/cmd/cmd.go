package cmd

import (
	"context"
	"fmt"
	"time"

	"petStore/internal/controller"

	"petStore/internal/middleware"

	// "github.com/goflyfox/gtoken"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// WebSocket
			s.BindHandler("/wsauto", func(r *ghttp.Request) {
				var ctx = r.Context()
				g.Log().Print(ctx, "WebSocket connected")
				ws, err := r.WebSocket()
				if err != nil {
					g.Log().Error(ctx, "WebSocket error:", err)
					r.Exit()
				}
				go func() {
					var lastTimeStamp = gtime.Now().TimestampMilli()
					for {
						if gtime.Now().TimestampMilli()-lastTimeStamp < 5000 {
							g.Log().Print(ctx, "current time:", gtime.Datetime(), gtime.Now().TimestampMilli()-lastTimeStamp)
							time.Sleep(1 * time.Second)
						} else {
							g.Log().Print(ctx, "Send time")
							if err = ws.WriteMessage(1, []byte(fmt.Sprint("Current time:", gtime.Datetime()))); err != nil {
								g.Log().Error(ctx, "WebSocket error:", err)
								return
							}
							lastTimeStamp = gtime.Now().TimestampMilli()
						}
					}
				}()
				for {
					msgType, msg, err := ws.ReadMessage()
					if err != nil {
						return
					}
					g.Log().Print(ctx, "Receive Message through Websocket, msg: ", msg)
					if err = ws.WriteMessage(msgType, []byte(fmt.Sprint("The server has received your message:", msg))); err != nil {
						return
					}
				}
			})

			s.BindHandler("/wschat", func(r *ghttp.Request) {
				controller.Chat.RequestWebSocket(r)
			})
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)
				// issue XTOKEN as the entry token
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Hello,
					)
					group.Group("/user", func(group *ghttp.RouterGroup) {
						// add XTOKEN to prevent csrf
						group.Middleware(middleware.NewCsrf())
						group.Bind(
							controller.User,
						)
					})
				})

				group.Group("/v2", func(group *ghttp.RouterGroup) {
					// group.Middleware(middleware.RefreshTokenAuth(), middleware.AccessTokenAuth())
					group.Middleware(middleware.AccessTokenAuth())
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
