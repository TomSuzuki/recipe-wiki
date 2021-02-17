package view

import "html/template"

// PageData ...全てのページを表示するときに必要な型のデータです。
type PageData struct {
	HTML      template.HTML
	PageTitle string
}

// TopPage ...トップページの表示に必要なデータです。
type TopPage struct {
}
