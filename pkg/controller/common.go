package controller

import "github.com/TomSuzuki/recipe-wiki/pkg/model"

// Controller ...コントローラーに渡すデータです。
type Controller struct {
	ServerData *model.ServerData
}

// Init ...コントローラーを初期化します。
func Init(sd *model.ServerData) Controller {
	return Controller{ServerData: sd}
}
