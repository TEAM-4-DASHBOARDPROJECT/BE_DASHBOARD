package route

import (
	"immersiveProject/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersiveProject/features/login/delivery"
	loginrepo "immersiveProject/features/login/repository"
	loginusecase"immersiveProject/features/login/usecase"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	/*  Route  
				*/
	
	e.POST("login", loginHandler.Login)
}
