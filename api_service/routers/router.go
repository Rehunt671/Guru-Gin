package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gitlab.com/gurugin/handlers"
)

func SetupRoutes(r fiber.Router, mlHandler handlers.MLHandler, recipeHander handlers.RecipeHandler) {
	// Serve Swagger documentation
	r.Get("/swagger/*", swagger.HandlerDefault)

	// Create a new group for versioned API routes
	v1 := r.Group("/v1")

	// Setup ML routes
	SetupMLRouter(v1, mlHandler)
	SetupRecipeRouter(v1, recipeHander)
}
