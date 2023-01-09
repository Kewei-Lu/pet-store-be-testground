package auth

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	INVALID_JWT_ERROR  = gerror.NewCode(gcode.New(30001, "invalid jwt token", "jwt is invalid"), "invalid jwt token")
	JWT_PAYLOAD_ERROR  = gerror.NewCode(gcode.New(30002, "error in asserting jwt payload", "payload is invalid"), "error in asserting jwt payload")
	JWT_GENERATE_ERROR = gerror.NewCode(gcode.New(30003, "error in generating jwt token", "jwt generating error"), "error in generating jwt token")
)
