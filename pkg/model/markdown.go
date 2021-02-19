package model

import (
	"io/ioutil"
)

// GetMarkdownText ...マークダウンを読み取って返す。
func GetMarkdownText(path string) (string, error) {
	md, err := ioutil.ReadFile(path)
	return string(md), err
}
