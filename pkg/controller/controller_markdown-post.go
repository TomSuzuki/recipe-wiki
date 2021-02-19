package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/recipe-wiki/pkg/config"
	"github.com/gin-gonic/gin"
)

// MarkdownPostController ...マークダウンページの保存を処理。
func (ctrl *Controller) MarkdownPostController(c *gin.Context) {
	// query[name]
	name, _ := service.QueryString(c, "name")
	if name != "curry" && name != "reference" {
		c.Status(http.StatusBadRequest)
		return
	}

	// form[text]
	text, b := c.GetPostForm("text")
	if !b {
		c.Status(http.StatusBadRequest)
		return
	}

	// path
	path := config.MediaFolder
	if name == "curry" {
		path += "curry.md"
	} else {
		path += "reference.md"
	}

	// save
	ioutil.WriteFile(path, []byte(text), 0666)
	c.Status(http.StatusOK)
}
