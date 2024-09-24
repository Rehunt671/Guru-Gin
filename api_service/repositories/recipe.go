package repositories

import (
	"gitlab.com/gurugin/models"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	FindRecipesByIngredients([]string) ([]models.Recipe, error)
}

type recipeRepository struct {
	db *gorm.DB
}

func NewMLRepository(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (r *recipeRepository) FindRecipesByIngredients(ingredientNames []string) ([]models.Recipe, error) {
	var recipes []models.Recipe

	err := r.db.Table("recipes").
		Select("recipes.*, users.*,menus.*").
		Joins("JOIN ingredients_on_recipes ON ingredients_on_recipes.recipe_id = recipes.id").
		Joins("JOIN ingredients ON ingredients.id = ingredients_on_recipes.ingredient_id").
		Joins("JOIN menus ON menus.id = recipes.menu_id").
		Joins("JOIN users ON users.id = recipes.user_id").
		Group("recipes.id, users.id, menus.id").
		Having("COUNT(DISTINCT ingredients.id) = COUNT(DISTINCT CASE WHEN ingredients.name IN (?) THEN ingredients.name ELSE NULL END)", ingredientNames).
		Find(&recipes).Error

	if err != nil {
		return nil, err
	}
	return recipes, nil
}
