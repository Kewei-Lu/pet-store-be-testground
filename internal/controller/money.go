package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"

	v1 "petStore/api/v1"
	"petStore/internal/model"
	"petStore/internal/service"
)

var (
	Money = cMoney{}
)

// User Handler
type cMoney struct{}

func (c *cMoney) GetMoney(ctx context.Context, req *v1.GetMoneyReq) (res *v1.GetMoneyRes, err error) {
	amount, err := service.UserMoney().QueryMoney(ctx, model.MoneyQueryInput{UserName: req.UserName})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(&v1.GetMoneyRes{Amount: amount})
	return
}

func (c *cMoney) AddMoney(ctx context.Context, req *v1.AddMoneyReq) (res *v1.AddMoneyRes, err error) {
	// check user validity from cookie
	UserNameCookie := g.RequestFromCtx(ctx).Cookie.Get("user-name").String()
	IssueTimeCookie := g.RequestFromCtx(ctx).Cookie.Get("issue-time").String()
	err = service.User().CookieValidate(ctx, model.UserCookiesInput{UserName: UserNameCookie, IssueTime: IssueTimeCookie})
	if err != nil {
		return nil, err
	}
	err = service.UserMoney().AddMoney2User(ctx, model.MoneyAddInput{UserName: req.SourceAccount, Amount: req.Amount, DestinationAccount: req.DestinationAccount, Comment: "Add fund"})
	if err != nil {
		var ReturnedError string = fmt.Sprintf("%d: %s", gerror.Code(err), err.Error())
		g.RequestFromCtx(ctx).Response.WriteStatus(403, &v1.AddMoneyRes{Success: false, Reason: ReturnedError})
		return
	}
	g.RequestFromCtx(ctx).Response.Writeln(&v1.UserRegistrationRes{Success: true, Reason: "success"})
	return
}

func (c *cMoney) TransferMoney(ctx context.Context, req *v1.TransferMoneyReq) (res *v1.TransferMoneyRes, err error) {
	// use jwt token to identify the client-user
	jwtToken := g.RequestFromCtx(ctx).Header.Get("jwt-token")
	g.Log().Print(ctx, "jwtToken: ", jwtToken)
	// decode
	token, err := jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Header)
		return JWT_SIG, nil
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(403, "token validation fails")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	g.Log().Print(ctx, "claims: ", claims)
	g.Log().Print(ctx, "sig: ", token.Signature)
	if ok {
		username := claims["user-name"].(string)
		if username != req.SourceAccount {
			g.RequestFromCtx(ctx).Response.WriteStatus(403, "invalid source user")
			return
		}
	}
	g.Log().Print(ctx, "parsed token: ", token)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(403)
		return
	}
	err = service.UserMoney().TransferMoney2User(ctx, model.MoneyTransferInput{UserName: req.SourceAccount, Amount: req.Amount, DestinationAccount: req.DestinationAccount, Comment: "Transfer Money"})
	if err != nil {
		var ReturnedError string = fmt.Sprintf("%d: %s", gerror.Code(err), err.Error())
		g.RequestFromCtx(ctx).Response.WriteStatus(403, &v1.AddMoneyRes{Success: false, Reason: ReturnedError})
		return
	}
	g.RequestFromCtx(ctx).Response.Writeln(&v1.UserRegistrationRes{Success: true, Reason: "success"})
	return
}
