package controller

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// TopPageController ...トップページの処理を行います。
func (ctrl *Controller) TopPageController(c *gin.Context) {

	// dto
	var data object.TopPageData
	data.RecipeDataList = ctrl.ServerData.TopRecipeList()

	// view
	view.NewView(c, object.PageData{
		HTML: view.TopPageView(data),
	})
}
