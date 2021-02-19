package model

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/TomSuzuki/recipe-wiki/pkg/config"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
)

// loadRecipe ...JSONファイルを読み込みます。
func loadRecipe() ServerData {
	bytes, _ := ioutil.ReadFile(config.RecipeFile)
	var recipeList ServerData
	json.Unmarshal(bytes, &recipeList)
	return recipeList
}

// saveRecipe ...変更を保存します。変更が完了したときのみ使用します。（通常 GET メソッドでの呼び出し時は使用しません。）
func (sd *ServerData) saveRecipe() error {
	bytes, _ := json.Marshal(sd)
	return ioutil.WriteFile(config.RecipeFile, bytes, 0666)
}

// ServerData ...とりあえず全てのデータを持ちます。
type ServerData struct {
	Recipe []object.RecipeData `json:"recipe"`
}

// Init ...検索データを初期化します。
func Init() *ServerData {
	os.Mkdir(config.MediaFolder, 0777)
	var serverData ServerData
	serverData = loadRecipe()
	return &serverData
}
