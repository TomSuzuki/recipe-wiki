package controller

import (
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// TopPageController ...トップページの処理を行います。
func (ctrl *Controller) TopPageController(c *gin.Context) {
	// query[keyword]
	keyword, _ := service.QueryString(c, "keyword")

	// dto
	var data object.TopPageData
	data.Keyword = keyword

	// list
	if keyword == "" {
		data.RecipeDataList = ctrl.ServerData.TopRecipeList()
	} else {
		data.RecipeDataList = ctrl.ServerData.RecipeListFilterKeyword(keyword)
	}

	// view
	view.NewView(c, object.PageData{
		HTML: view.TopPageView(data),
	})
}
