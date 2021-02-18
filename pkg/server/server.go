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
	// router.GET("/page/status", controller.PageStatusController)
	// router.GET("/error", controller.ErrorPageController)
	// router.GET("/page", controller.WordPageController)
	router.GET("/top", ctrl.TopPageController)
	// router.GET("/write", controller.WritePageController)
	// router.GET("/search", controller.SearchPageController)
	// router.GET("/edit", controller.EditPageController)

	// router.POST("/save", controller.SaveController)

	// router.DELETE("/page", controller.DeletePageController)

	// no route
	router.GET("", ctrl.TopPageController)
	// router.NoRoute(controller.ErrorPageController)

	return
}
