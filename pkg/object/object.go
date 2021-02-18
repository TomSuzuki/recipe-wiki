package object

import "html/template"

// PageData ...全てのページを表示するときに必要な型のデータです。
type PageData struct {
	HTML      template.HTML
	PageTitle string
}

// TopPageData ...トップページの表示に必要なデータです。
type TopPageData struct {
	RecipeDataList []RecipeData
}

// RecipePageData ...レシピページに必要なデータです。
type RecipePageData struct {
	RecipeData RecipeData
}

// WritePageData ...編集ページに必要なデータです。
type WritePageData struct {
	RecipeData RecipeData
}

// RecipeData ...レシピのデータです。
type RecipeData struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	IngredientList []Ingredient `json:"ingredients"`
	StepList       []string     `json:"steps"`
	MemoList       []string     `json:"memo"`
	Evaluation     int          `json:"evaluation"`
	ImagePath      string       `json:"image_path"`
}

// Ingredient ...材料のデータです。
type Ingredient struct {
	Name     string `json:"ingredient_name"`
	Quantity string `json:"ingredient_quantity"`
}
