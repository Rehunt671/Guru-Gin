package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/gurugin/services"
)

type RecipeHandler interface {
	FindRecipesByIngredients(*fiber.Ctx) error
}

type recipeHandler struct {
	mlService     services.MLService
	recipeService services.RecipeService
}

func NewRecipeHandler(mlService services.MLService, recipeService services.RecipeService) RecipeHandler {
	return &recipeHandler{
		mlService:     mlService,
		recipeService: recipeService,
	}
}

func (h *recipeHandler) FindRecipesByIngredients(ctx *fiber.Ctx) error {

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to get form data")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("No files uploaded")
	}

	ingredients, err := h.mlService.DetectObjectsWithGRPC(files)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to classify images: %v", err))
	}
	recipes, err := h.recipeService.FindRecipesByIngredients(ingredients)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to retrieve recipes: %v", err))
	}

	return ctx.JSON(fiber.Map{
		"recipes": recipes,
	})
}
