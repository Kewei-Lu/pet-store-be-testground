package money

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	INVALID_TOP_USER_ERROR                  = gerror.NewCode(gcode.New(20001, "invalid top-up UserName", "the user you are going to top-up is invalid"), "invalid top-up UserName")
	INVALID_TOP_AMOUNT_ERROR                = gerror.NewCode(gcode.New(20002, "invalid top-up Number", "the number you want to top-up is invalid"), "invalid top-up Number")
	INVALID_TRANSFER_SOURCE_USER_ERROR      = gerror.NewCode(gcode.New(20003, "invalid transfer source user", "the user you want to transfer is invalid"), "invalid transfer source user")
	INVALID_TRANSFER_DESTINATION_USER_ERROR = gerror.NewCode(gcode.New(20004, "invalid transfer destination user", "the user you want to transfer to is invalid"), "invalid transfer destination user")
	INVALID_TRANSFER_AMOUNT_ERROR           = gerror.NewCode(gcode.New(20005, "invalid transfer amount", "the amount you want to transfer is invalid"), "invalid transfer amount")
	INVALID_QUERY_USER_ERROR                = gerror.NewCode(gcode.New(20006, "invalid query UserName", "the user you are going to query is invalid"), "invalid query UserName")
	DUPLICATE_MONEY_ACCOUNT_ERROR           = gerror.NewCode(gcode.New(20007, "duplicated money account", "the user you are going to query is duplicated"), "duplicated money account")
)
