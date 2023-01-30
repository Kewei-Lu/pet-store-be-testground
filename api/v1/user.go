package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserRegistrationReq struct {
	g.Meta   `path:"/register" tags:"user" method:"post" summary:"User Registration api"`
	UserName string
	PassWord string
}
type UserRegistrationRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
	Reason  string
}

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"user" method:"post" summary:"User Login api"`
	UserName string
	PassWord string
}
type UserLoginRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
	Reason  string
}
type UserRegisterFromIntelReq struct {
	g.Meta   `path:"/registerFromIntel" tags:"user" method:"post" summary:"User Login api"`
	UserName string
	PassWord string
}
type UserRegisterFromIntelRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
	Reason  string
}
type UserCookiesReq struct {
	g.Meta    `path:"/cookies" tags:"user" method:"post" summary:"User Login api"`
	UserName  string
	IssueTime string
}
type UserCookiesRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
}
