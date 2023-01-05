package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TransferMoneyReq struct {
	g.Meta             `path:"/transfer" tags:"transition" method:"post" summary:"Money Transition api"`
	SourceAccount      string `p:"from"`
	DestinationAccount string `p:"to"`
	Amount             int
}
type TransferMoneyRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
	Reason  string
}

type AddMoneyReq struct {
	g.Meta             `path:"/add" tags:"transition" method:"post" summary:"Money Add api"`
	SourceAccount      string `p:"from"`
	DestinationAccount string `p:"to"`
	Amount             int
}
type AddMoneyRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Success bool
	Reason  string
}
type GetMoneyReq struct {
	g.Meta   `path:"/" tags:"transition" method:"get" summary:"Money Query api"`
	UserName string
}
type GetMoneyRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Amount int
}
