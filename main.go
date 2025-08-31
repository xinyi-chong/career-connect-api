package main

import (
	_ "gf_demo/internal/packed"

	_ "gf_demo/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
	// test push
}
