package user

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"time"

	"petStore/internal/model"
	"petStore/internal/service"
)

type sUser struct {
	users []*user // 本地内存存储所有已注册用户
}

type user struct {
	UserName    string
	Password    string
	LastLoginAt string
	RegisterAt  string
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{users: []*user{}}
}

func (s *sUser) CreateUser(ctx context.Context, in model.UserRegisterInput) error {
	for _, user := range s.users {
		if in.UserName == user.UserName {
			return DUPLICATED_USERNAME_ERROR
		}
	}
	s.users = append(s.users, &user{UserName: in.UserName, Password: in.PassWord, RegisterAt: string(time.Now().Unix()), LastLoginAt: "nil"})
	glog.Print(ctx, "current password store: ", s.users)
	return nil
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) error {
	for _, user := range s.users {
		if in.UserName == user.UserName {
			if user.Password == in.PassWord {
				user.LastLoginAt = time.Now().String()
				return nil
			}
		}
	}
	return WRONG_LOGIN_ERROR
}
