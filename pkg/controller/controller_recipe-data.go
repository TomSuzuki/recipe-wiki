package controller

import (
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/gin-gonic/gin"
)

// RecipeDataController ...レシピのデータをJSONで返す。
func (ctrl *Controller) RecipeDataController(c *gin.Context) {
	// query
	id, err := service.QueryInt(c, "id")
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}

	// view
	data, _ := ctrl.ServerData.GetRecipeData(id)
	c.JSON(http.StatusOK, data)
}
