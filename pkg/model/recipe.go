package model

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"time"

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

// SaveRecipeData ...レシピデータを保存する。
func (sd *ServerData) SaveRecipeData(data object.RecipeData) error {
	isNew := true
	for i := range sd.Recipe {
		if data.ID == sd.Recipe[i].ID {
			isNew = false
			if data.ImagePath == "" {
				data.ImagePath = sd.Recipe[i].ImagePath
			}
			sd.Recipe[i] = data
			break
		}
	}
	if isNew {
		sd.Recipe = append(sd.Recipe, data)
	}
	return sd.saveRecipe()
}

// CreateNewID ...新しいIDを発行する。
func (sd *ServerData) CreateNewID() (int, error) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1024; i++ {
		newID := rand.Intn(2*math.MaxInt32) + math.MinInt32
		if _, err := sd.GetRecipeData(newID); err != nil {
			return newID, nil
		}
	}
	return 0, errors.New("cannot create new id")
}

// RecipeListFilterKeyword ...キーワードに一致するレシピリストを返す。
func (sd *ServerData) RecipeListFilterKeyword(keyword string) []object.RecipeData {
	var hitName []object.RecipeData
	var hitIngredient []object.RecipeData

	for i := range sd.Recipe {
		r := sd.Recipe[i]
		if strings.Contains(r.Name, keyword) {
			hitName = append(hitName, r)
			continue
		}
		for j := range r.IngredientList {
			iName := r.IngredientList[j].Name
			if strings.Contains(iName, keyword) {
				hitIngredient = append(hitIngredient, r)
				break
			}
		}
	}

	return append(hitName, hitIngredient...)
}
