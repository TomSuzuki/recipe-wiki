package main

import "github.com/TomSuzuki/recipe-wiki/pkg/server"

func main() {
	server.InitLog()
	server.Router().Run(":8220")
}
