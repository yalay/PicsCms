package conf

import (
	"path/filepath"

	"github.com/unrolled/render"
)

var config Config

var Render = render.New(render.Options{
	Directory:  filepath.Join("views", "v3"),
	Extensions: []string{".tpl"},
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
