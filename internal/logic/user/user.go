package user

import (
	"context"
	"petStore/internal/model"

	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sUser) UserValidate(ctx context.Context, in model.UserValidateInput) error {
	for _, user := range s.users {
		if in.UserName == user.UserName {
			return nil
		}
	}
	return USER_VALIDATE_ERROR
}
func (s *sUser) UserSetLastLogin(ctx context.Context, in model.UserSetLastLoginInput) error {
	for _, user := range s.users {
		if in.UserName == user.UserName {
			user.LastLoginAt = gtime.TimestampMicroStr()
			return nil
		}
	}
	return USER_SET_LAST_LOGIN_ERROR
}
