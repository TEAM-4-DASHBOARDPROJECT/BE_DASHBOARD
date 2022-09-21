package route

import (
	"immersiveProject/config"
	"immersiveProject/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersiveProject/features/login/delivery"
	loginrepo "immersiveProject/features/login/repository"
	loginusecase "immersiveProject/features/login/usecase"

	userData "immersiveProject/features/users/data"
	userDelivery "immersiveProject/features/users/delivery"
	userUsecase "immersiveProject/features/users/usecase"

	classhandler "immersiveProject/features/class/delivery"
	classrepo "immersiveProject/features/class/repository"
	classusecase "immersiveProject/features/class/usecase"

	menteeData "immersiveProject/features/mentee/data"
	menteeDelivery "immersiveProject/features/mentee/delivery"
	menteeUsecase "immersiveProject/features/mentee/usecase"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	userData := userData.New(db)
	userUsecaseFactory := userUsecase.New(userData)
	userDelivery.New(e, userUsecaseFactory)

	classRepo := classrepo.New(db)
	classUsecase := classusecase.New(classRepo)
	classHandler := classhandler.New(classUsecase)

	menteeData := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeData)
	menteeDelivery.New(e, menteeUsecaseFactory)

	/*  Route
	 */

	e.POST("/login", loginHandler.Login)

	e.GET("/class", classHandler.GetClass, middlewares.JWTMiddleware())
	e.POST("/class", classHandler.Create, middlewares.JWTMiddleware())
	e.PUT("/class/:id", classHandler.Update, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", classHandler.Delete, middlewares.JWTMiddleware())

}
