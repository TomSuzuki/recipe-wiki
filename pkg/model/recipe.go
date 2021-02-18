package model

import (
	"errors"
	"math"

	"github.com/TomSuzuki/recipe-wiki/pkg/object"
)

// TopRecipeList ...トップページに表示するレシピのリストを生成します。
func (sd *ServerData) TopRecipeList() []object.RecipeData {
	return sd.Recipe[:int(math.Min(32, float64(len(sd.Recipe))))] // あとでランダムにする。
}

// GetRecipeData ...idを指定してレシピデータを取得する。
func (sd *ServerData) GetRecipeData(id int) (object.RecipeData, error) {
	for i := range sd.Recipe {
		if id == sd.Recipe[i].ID {
			return sd.Recipe[i], nil
		}
	}
	return object.RecipeData{}, errors.New("recipe data is not found")
}
