package controller

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/TomSuzuki/recipe-wiki/pkg/config"
	"github.com/TomSuzuki/recipe-wiki/pkg/object"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

// RecipePostController ...レシピのデータを保存する。
func (ctrl *Controller) RecipePostController(c *gin.Context) {
	// receive data
	var recipePostData object.RecipePostData
	c.BindJSON(&recipePostData)

	// data
	var data object.RecipeData

	// id check
	if recipePostData.ID == nil {
		data.ID, _ = ctrl.ServerData.CreateNewID()
	} else {
		data.ID = *recipePostData.ID
	}

	// set data
	data.Name = recipePostData.Name
	data.IngredientList = recipePostData.IngredientList
	data.StepList = recipePostData.StepList
	data.MemoList = recipePostData.MemoList
	data.Evaluation = recipePostData.Evaluation

	// image
	if recipePostData.ImageData != "" {
		// decode
		temp, _ := url.QueryUnescape(recipePostData.ImageData)
		temp = strings.Split(temp, ",")[1]
		byteData, _ := base64.StdEncoding.DecodeString(temp)
		imageData, _, err := image.Decode(bytes.NewReader(byteData))
		if err == nil {
			// resize
			size := int(math.Min(float64(imageData.Bounds().Dx()), float64(imageData.Bounds().Dy())))
			imageData = imaging.Fill(imageData, size, size, imaging.Center, imaging.Lanczos)
			buffer := new(bytes.Buffer)
			jpeg.Encode(buffer, imageData, nil)
			imgJPEG, _, _ := image.Decode(bytes.NewReader(buffer.Bytes()))
			thumbnail := resize.Thumbnail(uint(256), uint(256), imgJPEG, resize.Lanczos3)
			dataJPEG := new(bytes.Buffer)
			jpeg.Encode(dataJPEG, thumbnail, &jpeg.Options{Quality: 80})
			imgJPEGBytes := dataJPEG.Bytes()

			// save
			os.Mkdir(config.MediaFolder, 0777)
			h := sha1.Sum(imgJPEGBytes)
			fname := fmt.Sprintf("%s_%x.jpg", time.Now().Format("20060102150405"), h[:16])
			fpath := config.MediaFolder + fname
			ioutil.WriteFile(fpath, imgJPEGBytes, 0600)

			// path
			data.ImagePath = fname
		}
	}

	// set data
	ctrl.ServerData.SaveRecipeData(data)

	// response
	resData, _ := ctrl.ServerData.GetRecipeData(data.ID)
	c.JSON(http.StatusOK, resData)
}
