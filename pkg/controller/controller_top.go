package controller

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// TopPageController ...トップページの処理を行います。
func (ctrl *Controller) TopPageController(c *gin.Context) {

	// get recipe list
	ctrl.ServerData.TopRecipeList()

	// dto
	var data object.TopPage

	// view
	view.NewView(c, object.PageData{
		HTML: view.TopPageView(data),
	})
}
