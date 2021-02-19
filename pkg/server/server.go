package server

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/controller"
	"github.com/gin-gonic/gin"
)

// Router ...ルーティングを行います。
func Router(ctrl controller.Controller) (router *gin.Engine) {

	// root
	router = gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")
	router.Static("/data", "./data")

	// route
	router.GET("/top", ctrl.TopPageController)

	router.GET("/recipe", ctrl.RecipePageController)
	router.GET("/recipe/data", ctrl.RecipeDataController)
	router.POST("/recipe/data", ctrl.RecipePostController)

	router.GET("/write", ctrl.WritePageController)

	router.GET("/markdown", ctrl.MarkdownPageController)

	// no route
	router.GET("", ctrl.TopPageController)
	//router.NoRoute(ctrl.ErrorPageController)

	return
}
