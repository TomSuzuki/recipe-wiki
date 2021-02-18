package model

import (
	"math"

	"github.com/TomSuzuki/recipe-wiki/pkg/object"
)

// TopRecipeList ...トップページに表示するレシピのリストを生成します。
func (sd *ServerData) TopRecipeList() []object.RecipeData {
	return sd.Recipe[:int(math.Min(5, float64(len(sd.Recipe))))]
}
