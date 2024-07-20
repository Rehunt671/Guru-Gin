package routers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/handlers"
)

func SetupMLRouter(router fiber.Router, mlHandler handlers.MLHandler) {
	ml := router.Group("/ml")
	ml.Post("/train", mlHandler.TrainModel)
	ml.Post("/detectObject", mlHandler.DetectObjects)
}
