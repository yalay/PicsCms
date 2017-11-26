package conf

import (
	"fmt"
	"html/template"
	"models"
	"path/filepath"

	"github.com/BurntSushi/toml"
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
			"tagUrl":     GenTagUrl,
		},
	},
})

type Config struct {
	WebUrl            string
	WebName           string
	WebKeywords       string
	WebDesc           string
	RootPath          string
	AttachProfileName string
	Cates             []*models.Category
}

func init() {
	_, err := toml.DecodeFile("sys.profile", &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}

func WebUrl() string {
	return config.WebUrl
}

func WebName() string {
	return config.WebName
}

func WebKeywords() string {
	return config.WebKeywords
}

func WebDesc() string {
	return config.WebDesc
}

func RootPath() string {
	return config.RootPath
}

func ArticleProfileName() string {
	return config.AttachProfileName
}

func TotalCates() []*models.Category {
	return config.Cates
}
