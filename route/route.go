package route

import (
	"immersiveProject/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginhandler "immersiveProject/features/login/delivery"
	loginrepo "immersiveProject/features/login/repository"
	loginusecase "immersiveProject/features/login/usecase"

	userData "immersiveProject/features/users/data"
	userDelivery "immersiveProject/features/users/delivery"
	userUsecase "immersiveProject/features/users/usecase"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, cfg *config.AppConfig) {
	loginRepo := loginrepo.New(db)
	loginUsecase := loginusecase.New(loginRepo)
	loginHandler := loginhandler.New(loginUsecase)

	userData := userData.New(db)
	userUsecaseFactory := userUsecase.New(userData)
	userDelivery.New(e, userUsecaseFactory)

	/*  Route
	 */

	e.POST("login", loginHandler.Login)

}
