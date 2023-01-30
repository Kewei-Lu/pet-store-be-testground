package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

// type AccessTokenReq struct {
// 	g.Meta       `path:"/access-token" tags:"Hello" method:"post" summary:"Require an access-token for user"`
// 	RefreshToken string
// }
// type AccessTokenRes struct {
// 	g.Meta      `mime:"text/html" example:"string"`
// 	AccessToken string
// }

type AuthReLoginReq struct {
	g.Meta `path:"/relogin" tags:"auth" method:"post" summary:"User Login api"`
	Token  string
}
type AuthReLoginRes struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserName string
	Success  bool
}
