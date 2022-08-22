package main

import (
	_ "poetry/internal/packed"

	_ "poetry/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"poetry/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
