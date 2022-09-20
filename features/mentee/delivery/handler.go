package delivery

import (
	"immersiveProject/features/mentee"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

// func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
// 	return &MenteeDelivery{
// 		menteeUsecase: usecase,
// 	}
// }
