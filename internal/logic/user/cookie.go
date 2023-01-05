package user

import (
	"context"
	"petStore/internal/model"
	"strconv"
	"time"
)

func (s *sUser) CookieValidate(ctx context.Context, in model.UserCookiesInput) error {
	for _, user := range s.users {
		if in.UserName == user.UserName {
			IssueTime, err := strconv.ParseInt(in.IssueTime, 10, 64)
			if err != nil {
				return err
			}
			if IssueTime-time.Now().Unix() <= 60*60*24*365 {
				user.LastLoginAt = string(time.Now().Unix())
				return nil
			} else {
				return COOKIE_EXPIRE_ERROR
			}
		}
	}
	return COOKIE_UNKNOWN_ERROR
}
