package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"petStore/api/v1"
	"petStore/internal/model"
	"petStore/internal/service"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

// for login with X-Token and jwt checked
func (c *cAuth) ReLogin(ctx context.Context, req *v1.AuthReLoginReq) (res *v1.AuthReLoginRes, err error) {
	g.Log().Print(ctx, "Receive token:", req.Token)
	// check jwt token
	err = service.Auth().ValidateJwtToken(ctx, model.JwtValidateInput{Token: req.Token, Sig: JWT_SIG})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// query login username for jwt
	queryRes, err := service.Auth().QueryJwtToken(ctx, model.JwtQueryInput{Token: req.Token, Sig: JWT_SIG})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// validate username is valid
	err = service.User().UserValidate(ctx, model.UserValidateInput{UserName: queryRes.UserName})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	// set last login time
	err = service.User().UserSetLastLogin(ctx, model.UserSetLastLoginInput{LastLoginTimeStamp: gtime.Now().TimestampMicroStr(), UserName: queryRes.UserName})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.AuthReLoginRes{UserName: "nil", Success: false})
		return
	}
	g.RequestFromCtx(ctx).Response.Writeln(&v1.AuthReLoginRes{Success: true, UserName: queryRes.UserName})
	return
}
