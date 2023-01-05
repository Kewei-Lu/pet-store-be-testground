package main

import (
	_ "petStore/internal/logic"
	_ "petStore/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"petStore/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
