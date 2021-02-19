package controller

import (
	"html/template"
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/recipe-wiki/pkg/config"
	"github.com/TomSuzuki/recipe-wiki/pkg/model"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/TomSuzuki/recipe-wiki/pkg/view"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// MarkdownPageController ...マークダウンページの表示を処理。
func (ctrl *Controller) MarkdownPageController(c *gin.Context) {
	// query [name]
	name, _ := service.QueryString(c, "name")
	if name != "curry" && name != "reference" {
		c.Status(http.StatusBadRequest) // エラーページの表示に変更する。
		return
	}

	// query [edit]
	edit, _ := service.QueryInt(c, "edit")
	isEdit := edit == 1

	// data
	var data object.MarkdownPageData
	path := config.MediaFolder
	if name == "curry" {
		path += "curry.md"
		data.MarkdownTitle = "カレー"
	} else {
		path += "reference.md"
		data.MarkdownTitle = "参考サイト"
	}

	// markdown
	data.MarkdownText, _ = model.GetMarkdownText(path)
	data.MarkdownHTML = template.HTML(string(blackfriday.MarkdownCommon([]byte(data.MarkdownText))))
	data.MarkdownEdit = isEdit

	// view
	view.NewView(c, object.PageData{
		HTML: view.MarkdownPageView(data),
	})
}
