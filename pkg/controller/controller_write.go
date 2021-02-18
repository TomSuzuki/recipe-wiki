package controller

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// WritePageController ...編集ページの処理。
func (ctrl *Controller) WritePageController(c *gin.Context) {
	// dto
	var data object.WritePageData

	// view
	view.NewView(c, object.PageData{
		HTML: view.WritePageView(data),
	})
}
