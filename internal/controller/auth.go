package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v2 "petStore/api/v2"
	"petStore/internal/model"
	"petStore/internal/service"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

// relogin only requires refresh token
func (c *cAuth) ReLogin(ctx context.Context, req *v2.AuthReLoginReq) (res *v2.AuthReLoginRes, err error) {
	g.Log().Print(ctx, "Receive token:", req.Token)
	// check refresh token
	err = service.Auth().ValidateJWTToken(ctx, model.JWTValidateInput{Token: req.Token, Sig: REFRESH_TOKEN_SIG})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v2.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// query login username for refresh token
	queryRes, err := service.Auth().QueryRefreshToken(ctx, model.RefreshTokenQueryInput{Token: req.Token, RefreshTokenSig: REFRESH_TOKEN_SIG})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v2.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// validate username is valid
	err = service.User().UserValidate(ctx, model.UserValidateInput{UserName: queryRes.UserName})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v2.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// set last login time
	err = service.User().UserSetLastLogin(ctx, model.UserSetLastLoginInput{LastLoginTimeStamp: gtime.Now().TimestampMicroStr(), UserName: queryRes.UserName})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v2.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// issue access token
	AccessToken, err := service.Auth().IssueAccessToken(ctx, model.AccessTokenIssueInput{UserName: queryRes.UserName, IssueTime: gtime.TimestampMicroStr(), AccessTokenSig: ACCESS_TOKEN_SIG})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, "error in generating access-token")
		return
	}
	g.RequestFromCtx(ctx).Cookie.SetCookie("access-token", AccessToken.Token, "", "/", gtime.H*12)
	g.RequestFromCtx(ctx).Response.Writeln(&v2.AuthReLoginRes{Success: true, UserName: queryRes.UserName})
	return
}
