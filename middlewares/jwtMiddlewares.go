package middlewares

import (
	"fmt"
	"immersiveProject/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.SECRET_JWT),
	})
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId 
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(c echo.Context) (int, error) {
	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_JWT), nil
	})

	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId),nil
	}

	return -1, fmt.Errorf("token invalid")
}
