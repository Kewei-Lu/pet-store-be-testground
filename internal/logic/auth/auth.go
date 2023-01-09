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

func (s *sAuth) ValidateJwtToken(ctx context.Context, in model.JwtValidateInput) error {
	g.Log().Print(ctx, "jwtToken: ", in.Token)
	_, err := jwt.ParseWithClaims(in.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SIG, nil
	})
	if err != nil {
		return INVALID_JWT_ERROR
	}
	return nil
}

func (s *sAuth) QueryJwtToken(ctx context.Context, in model.JwtQueryInput) (*model.JwtQueryOutput, error) {
	g.Log().Print(ctx, "jwtToken: ", in.Token)
	token, err := jwt.ParseWithClaims(in.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Header)
		return in.Sig, nil
	})
	if err != nil {
		return nil, INVALID_JWT_ERROR
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	g.Log().Print(ctx, "claims: ", claims)
	if ok {
		username := claims["user-name"].(string)
		issueTime := claims["issue-time"].(string)
		return &model.JwtQueryOutput{UserName: username, IssueTime: issueTime}, nil
	}
	return nil, JWT_PAYLOAD_ERROR
}

func (s *sAuth) IssueJwtToken(ctx context.Context, in model.JwtIssueInput) (*model.JwtIssueOutPut, error) {
	g.Log().Print(ctx, "Issue JWT token for user: ", in.UserName)
	// generating jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user-name": in.UserName, "issue-time": in.IssueTime})
	signingString, err := claims.SignedString(in.Sig)
	if err != nil {
		return nil, JWT_GENERATE_ERROR
	}
	g.Log().Print(ctx, "signingSting: ", signingString)
	return &model.JwtIssueOutPut{Token: signingString}, nil
}
