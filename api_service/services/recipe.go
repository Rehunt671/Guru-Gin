package services

import (
	"gitlab.com/gurugin/models"
	"gitlab.com/gurugin/repositories"
)

type RecipeService interface {
	FindRecipesByIngredients(ingredients []string) ([]models.Recipe, error)
}

type recipeService struct {
	mlService  MLService
	recipeRepo repositories.RecipeRepository
}

func NewRecipeService(mlService MLService, recipeRepo repositories.RecipeRepository) RecipeService {
	return &recipeService{
		mlService:  mlService,
		recipeRepo: recipeRepo,
	}
}

func (s *recipeService) FindRecipesByIngredients(ingredients []string) ([]models.Recipe, error) {
	return s.recipeRepo.FindRecipesByIngredients(ingredients)
}
