package auth

import (
	"context"
	"fmt"
	"petStore/internal/model"
	"petStore/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

func (s *sAuth) ValidateJWTToken(ctx context.Context, in model.JWTValidateInput) error {
	g.Log().Print(ctx, "JWT Token: ", in.Token)
	_, err := jwt.ParseWithClaims(in.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return in.Sig, nil
	})
	if err != nil {
		return INVALID_JWT_ERROR
	}
	return nil
}

func (s *sAuth) QueryAccessToken(ctx context.Context, in model.AccessTokenQueryInput) (*model.AccessTokenQueryOutput, error) {
	g.Log().Print(ctx, "Access Token: ", in.Token)
	token, err := jwt.ParseWithClaims(in.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Header)
		return in.AccessTokenSig, nil
	})
	if err != nil {
		return nil, INVALID_JWT_ERROR
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	g.Log().Print(ctx, "claims: ", claims)
	if ok {
		username := claims["user-name"].(string)
		issueTime := claims["issue-time"].(string)
		tokenType := claims["type"].(string)
		if tokenType != "access-token" {
			return nil, INVALID_JWT_ERROR
		}
		return &model.AccessTokenQueryOutput{UserName: username, IssueTime: issueTime}, nil
	}
	return nil, JWT_PAYLOAD_ERROR
}

func (s *sAuth) QueryRefreshToken(ctx context.Context, in model.RefreshTokenQueryInput) (*model.RefreshTokenQueryOutput, error) {
	g.Log().Print(ctx, "Refresh Token: ", in.Token)
	token, err := jwt.ParseWithClaims(in.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Header)
		return in.RefreshTokenSig, nil
	})
	if err != nil {
		return nil, INVALID_JWT_ERROR
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	g.Log().Print(ctx, "claims: ", claims)
	if ok {
		username := claims["user-name"].(string)
		issueTime := claims["issue-time"].(string)
		tokenType := claims["type"].(string)
		if tokenType != "refresh-token" {
			return nil, INVALID_JWT_ERROR
		}
		return &model.RefreshTokenQueryOutput{UserName: username, IssueTime: issueTime}, nil
	}
	return nil, JWT_PAYLOAD_ERROR
}
