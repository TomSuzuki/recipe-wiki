package view

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/TomSuzuki/recipe-wiki/pkg/config"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/gin-gonic/gin"
)

// NewView ...表示する。
func NewView(c *gin.Context, data object.PageData) {
	if data.PageTitle == "" {
		data.PageTitle = config.ServiceName
	}
	c.HTML(http.StatusOK, "index.html", data)
}

// createHTML ...htmlの生成。
func createHTML(file string, data interface{}) template.HTML {
	var body bytes.Buffer
	t := template.Must(template.ParseFiles(file))
	t.Execute(&body, data)
	return template.HTML(body.String())
}

// TopPageView ...トップページの表示。
func TopPageView(data object.TopPageData) template.HTML {
	return createHTML("templates/top.html", data)
}

// RecipePageView ...レシピページの表示。
func RecipePageView(data object.RecipePageData) template.HTML {
	return createHTML("templates/recipe.html", data)
}
