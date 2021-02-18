package main

import (
	"github.com/TomSuzuki/recipe-wiki/pkg/controller"
	"github.com/TomSuzuki/recipe-wiki/pkg/model"
	"github.com/TomSuzuki/recipe-wiki/pkg/server"
)

func main() {
	sd := model.Init()
	ctrl := controller.Init(sd)
	server.InitLog()
	server.Router(ctrl).Run(":8220")
}
