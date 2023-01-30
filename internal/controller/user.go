package controller

import (
	"context"
	"fmt"
	"time"

	v1 "petStore/api/v1"
	"petStore/internal/model"
	"petStore/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	User = cUser{}
)

// User Handler
type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *v1.UserRegistrationReq) (res *v1.UserRegistrationRes, err error) {
	err = service.User().CreateUser(ctx, model.UserRegisterInput{UserName: req.UserName, PassWord: req.PassWord})
	if err != nil {
		var ReturnedError string = fmt.Sprintf("%d: %s", gerror.Code(err), err.Error())
		g.RequestFromCtx(ctx).Response.WriteStatus(403, &v1.UserRegistrationRes{Success: false, Reason: ReturnedError})
		return
	}
	// add money account
	err = service.UserMoney().CreateMoneyAccount(ctx, model.CreateMoneyAccountInput{UserName: req.UserName, CreationTimeStamp: gtime.TimestampMilliStr()})
	// add 2000 for default
	err = service.UserMoney().AddMoney2User(ctx, model.MoneyAddInput{UserName: req.UserName, DestinationAccount: req.UserName, Amount: 2000000, Comment: "fund for registration"})
	g.RequestFromCtx(ctx).Response.Writeln(&v1.UserRegistrationRes{Success: true, Reason: "success"})
	return
}

func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	glog.Print(ctx, "user", req.UserName, "is going to login")
	err = service.User().Login(ctx, model.UserLoginInput{UserName: req.UserName, PassWord: req.PassWord})
	if err != nil {
		var ReturnedError string = fmt.Sprintf("%d: %s", gerror.Code(err).Code(), err.Error())
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.UserLoginRes{Success: false, Reason: ReturnedError})
		return
	}
	// generate refresh token
	RefreshToken, err := service.Auth().IssueRefreshToken(ctx, model.RefreshTokenIssueInput{UserName: req.UserName, RefreshTokenSig: REFRESH_TOKEN_SIG, IssueTime: gtime.TimestampMicroStr()})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, "fail to generate refresh-token")
		return
	}
	// generate access token
	AccessToken, err := service.Auth().IssueAccessToken(ctx, model.AccessTokenIssueInput{UserName: req.UserName, AccessTokenSig: ACCESS_TOKEN_SIG, IssueTime: gtime.TimestampMicroStr()})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(401, "fail to generate access-token")
		return
	}
	// set token stored in cookies
	g.RequestFromCtx(ctx).Cookie.SetCookie("issue-time", fmt.Sprint(time.Now().Unix()), "", "/", gtime.D*365)
	g.RequestFromCtx(ctx).Cookie.SetCookie("user-name", req.UserName, "", "/", gtime.D*365)
	g.RequestFromCtx(ctx).Cookie.SetCookie("refresh-token", RefreshToken.Token, "", "/", gtime.D*7)
	g.RequestFromCtx(ctx).Cookie.SetCookie("access-token", AccessToken.Token, "", "/", gtime.H*12)
	g.RequestFromCtx(ctx).Response.Writeln(&v1.UserLoginRes{Success: true, Reason: "success"})
	return
}

func (c *cUser) CheckCookies(ctx context.Context, req *v1.UserCookiesReq) (res *v1.UserCookiesRes, err error) {
	err = service.User().CookieValidate(ctx, model.UserCookiesInput{UserName: req.UserName, IssueTime: req.IssueTime})
	if err != nil {
		var ReturnedError string = fmt.Sprintf("%d: %s", gerror.Code(err).Code(), err.Error())
		g.RequestFromCtx(ctx).Response.WriteStatus(401, &v1.UserLoginRes{Success: false, Reason: ReturnedError})
		glog.Print(ctx, "res: ", res)
		return
	}
	g.RequestFromCtx(ctx).Response.Writeln(&v1.UserLoginRes{Success: true, Reason: "success"})
	glog.Print(ctx, "res: ", res)

	return
}
