package auth

import (
	"context"
	"petStore/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
)

func (s *sAuth) IssueAccessToken(ctx context.Context, in model.AccessTokenIssueInput) (*model.AccessTokenIssueOutput, error) {
	g.Log().Print(ctx, "Issue Access token for user: ", in.UserName)
	// generating access token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user-name": in.UserName, "issue-time": in.IssueTime, "type": "access-token"})
	signingString, err := claims.SignedString(in.AccessTokenSig)
	if err != nil {
		return nil, JWT_GENERATE_ERROR
	}
	g.Log().Print(ctx, "signingSting: ", signingString)
	return &model.AccessTokenIssueOutput{Token: signingString}, nil
}

func (s *sAuth) IssueRefreshToken(ctx context.Context, in model.RefreshTokenIssueInput) (*model.RefreshTokenIssueOutput, error) {
	g.Log().Print(ctx, "Issue Refresh token for user: ", in.UserName)
	// generating refresh token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user-name": in.UserName, "issue-time": in.IssueTime, "type": "refresh-token"})
	signingString, err := claims.SignedString(in.RefreshTokenSig)
	if err != nil {
		return nil, JWT_GENERATE_ERROR
	}
	g.Log().Print(ctx, "signingSting: ", signingString)
	return &model.RefreshTokenIssueOutput{Token: signingString}, nil
}
