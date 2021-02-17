package controller

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
)

// TopPageController ...トップページの処理を行います。
func TopPageController(c *gin.Context) {

	// view
	view.NewView(c, view.PageData{
		HTML: view.TopPageView(view.TopPage{}),
	})
}
