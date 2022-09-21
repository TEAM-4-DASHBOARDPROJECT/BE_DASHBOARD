package main

import (

	"immersiveProject/config"
	"immersiveProject/utils/database/mysql"
	"immersiveProject/route"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	route.InitRoutes(e, db, cfg)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}