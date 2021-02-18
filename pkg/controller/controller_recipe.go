package controller

import (
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// RecipePageController ...レシピページの生成。
func (ctrl *Controller) RecipePageController(c *gin.Context) {
	// query
	id, err := service.QueryInt(c, "id")
	if err != nil {
		// エラーページの表示
		return
	}

	// dto
	var data object.RecipePageData
	data.RecipeData, _ = ctrl.ServerData.GetRecipeData(id)

	// view
	view.NewView(c, object.PageData{
		HTML: view.RecipePageView(data),
	})
}
