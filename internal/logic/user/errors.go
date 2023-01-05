package user

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	DUPLICATED_USERNAME_ERROR = gerror.NewCode(gcode.New(10001, "duplicated UserName", "the user name already exists when registering"), "UserName already exists")
	WRONG_LOGIN_ERROR         = gerror.NewCode(gcode.New(10002, "login information not correct", "use the wrong login information to log in"), "login information wrong")
	COOKIE_EXPIRE_ERROR       = gerror.NewCode(gcode.New(10003, "cookies expires", "the cookies have expired"), "the cookies have expired")
	COOKIE_UNKNOWN_ERROR      = gerror.NewCode(gcode.New(10004, "unknown cookie user", "the cookies user is unknown"), "the cookies user is unknown")
)
