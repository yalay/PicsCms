package conf

import (
	"html/template"
	"path/filepath"

	"github.com/unrolled/render"
)

var config Config

var Render = render.New(render.Options{
	Directory:  filepath.Join("views", "v3"),
	Extensions: []string{".tpl"},
	Funcs: []template.FuncMap{
		{
			"articleUrl": GenArticleUrl,
			"attachUrl":  GenAttachUrl,
			"cateUrl":    GenCateUrl,
		},
	},
})

type Config struct {
	RootPath    string
	ProfileName string
}

func RootPath() string {
	return "test"
	//return config.RootPath
}

func ArticleProfileName() string {
	return ".profile"
}

func CateProfilePath() string {
	return "cate.profile"
}
