package controllers

import (
	"back-end/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type RecipeController struct{}

func (_ *RecipeController) GetRecipes(c *gin.Context) {
	var recipe model.Recipe
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		log.Fatalf("Page must be integer")
	}
	recipes, err := recipe.GetByPage(page)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, recipes)

}
func (_ *RecipeController) GetRecipe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	recipe, err := (&model.Recipe{}).GetByID(id)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, recipe)

}
