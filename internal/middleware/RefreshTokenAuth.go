package middleware

import (
	"fmt"
	"net/http"
	"petStore/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
)

var (
	RefreshTokenName string = "refresh-token"
)

func RefreshTokenAuth() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		// Read the token attached to the request
		// Read priority: Router < Query < Body < Form < Custom < Header
		RefreshToken := r.Header.Get(RefreshTokenName)
		if RefreshToken == "" {
			r.Response.WriteStatusExit(http.StatusForbidden, "Fail to fetch refresh token")
			return
		}
		token, err := jwt.ParseWithClaims(RefreshToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			fmt.Println(token.Header)
			return controller.REFRESH_TOKEN_SIG, nil
		})
		if err != nil {
			r.Response.WriteStatusExit(http.StatusForbidden, "Invalid refresh token")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			username := claims["user-name"].(string)
			r.Request.Header.Add("RefreshTokenDecodedUserName", username)
		} else {
			r.Response.WriteStatusExit(http.StatusForbidden, "Fail to fetch user name")
			return
		}
		r.Middleware.Next()
	}
}
