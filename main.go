package main

import (
	"fmt"

	"immersiveProject/config"
	"immersiveProject/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e.Pre(middleware.RemoveTrailingSlash())

	routes.InitRoutes(e, db, cfg)

	err := e.Start(":" + cfg.SERVER_PORT)

	if err != nil {
		panic(err)
	}
}
