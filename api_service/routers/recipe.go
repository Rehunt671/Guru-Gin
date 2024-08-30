package routers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/handlers"
)

func SetupRecipeRouter(router fiber.Router, recipeHander handlers.RecipeHandler) {
	recipe := router.Group("/recipe")
	recipe.Post("/findRecipeOnIngredients", recipeHander.FindRecipesByIngredients)
}
