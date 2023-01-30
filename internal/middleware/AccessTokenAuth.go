package middleware

import (
	"errors"
	"net/http"
	"petStore/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
)

var (
	AccessTokenName  string = "access-token"
	UserNameQueryErr error  = errors.New("error in UserNameQuery")
)

func UserNameQueryFromJWT(jwtToken, userNameFieldText string, Sig []byte) (string, error) {
	token, err := jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Sig, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		username := claims[userNameFieldText].(string)
		return username, nil
	} else {
		return "", UserNameQueryErr
	}
}

func AccessTokenAuth() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		// Read the token attached to the request
		// Read priority: Router < Query < Body < Form < Custom < Header
		AccessToken := r.Header.Get(AccessTokenName)
		if AccessToken == "" {
			r.Response.WriteStatusExit(http.StatusForbidden, "Fail to fetch access token")
			return
		}
		UserNameInAccessToken, err := UserNameQueryFromJWT(AccessToken, "user-name", controller.ACCESS_TOKEN_SIG)
		if err != nil {
			r.Response.WriteStatusExit(http.StatusForbidden, "Fail to parse user-name from access token")
			return
		}
		// when using access token, a refresh token must exist in cookie
		// we must assure user-name inside them are equal
		refreshToken, err := r.Request.Cookie(RefreshTokenName)
		if err != nil {
			r.Response.WriteStatusExit(http.StatusForbidden, "Refresh Token fetch error when validating access token")
			return
		}
		UserNameInRefreshToken, err := UserNameQueryFromJWT(refreshToken.Value, "user-name", controller.REFRESH_TOKEN_SIG)
		if err != nil {
			r.Response.WriteStatusExit(http.StatusForbidden, "Fail to parse user-name from refresh token")
			return
		}
		if UserNameInRefreshToken != UserNameInAccessToken {
			r.Response.WriteStatusExit(http.StatusForbidden, "access token and refresh token validation fail")
			return
		}
		r.Middleware.Next()
	}
}
