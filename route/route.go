package route

import (
	"immersiveProject/config"
	"immersiveProject/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersiveProject/features/login/delivery"
	loginrepo "immersiveProject/features/login/repository"
	loginusecase "immersiveProject/features/login/usecase"

	classhandler "immersiveProject/features/class/delivery"
	classrepo "immersiveProject/features/class/repository"
	classusecase "immersiveProject/features/class/usecase"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	classRepo := classrepo.New(db)
	classUsecase := classusecase.New(classRepo)
	classHandler := classhandler.New(classUsecase)

	/*  Route  
				*/
	
	e.POST("/login", loginHandler.Login)

	e.GET("/class", classHandler.GetClass, middlewares.JWTMiddleware())
	e.POST("/class", classHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/class", classHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/class", classHandler.Delete, middlewares.JWTMiddleware())
}
