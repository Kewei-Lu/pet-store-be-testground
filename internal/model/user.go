package model

// UserRegisterInput 用户注册输入
type UserRegisterInput struct {
	UserName string // 注册用户名
	PassWord string // 注册密码
}

type UserLoginInput struct {
	UserName string // 登录用户名
	PassWord string // 登录密码
}

type UserValidateInput struct {
	UserName string // 登录用户名
}
type UserSetLastLoginInput struct {
	UserName           string
	LastLoginTimeStamp string
}
type UserCookiesInput struct {
	UserName  string // user-name from cookie `p:user-name`
	IssueTime string // issue-time from cookie `p:issue-time`
}
