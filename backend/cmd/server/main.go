package main

import (
	"fmt"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/router"
)

func main() {
	r := router.InitRouter()
	r.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port)) // 后端端口
}
