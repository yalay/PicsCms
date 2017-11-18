package main

import (
	"conf"
	"controllers"
	"flag"
	"net/http"
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/file"
)

var (
	port  int
	debug bool
)

func init() {
	flag.IntVar(&port, "p", 8000, "listen port")
	flag.BoolVar(&debug, "d", false, "whether debug")
	flag.Parse()

	conf.InitLog("./test.log", debug)
}

func main() {
	router := routing.New()
	router.Get(`/`, controllers.HomeHandler)
	router.Get(`/<cate:(bigbreast|naked|hotass|bras)>.html`, controllers.CateHandler)
	router.Get(`/<cate:(bigbreast|naked|hotass|bras)>-<pid:[pn\d]+>.html`, controllers.CateHandler)
	router.Get(`/article-<id:\d+>.html`, controllers.ArticleHandler)
	router.Get(`/article-<id:\d+>-<pid:[pn\d]+>.html`, controllers.ArticleHandler)
	router.Get(`/tags-<tag:[^(.html)\s]+>.html`, controllers.TagsHandler)
	router.Get(`/tags-<tag:[^-\s]+>-<pid:[pn\d]+>.html`, controllers.TagsHandler)

	// static file
	router.Get("/*", file.Server(file.PathMap{
		"/css":     "./views/v3/css",
		"/img":     "./views/v3/img",
		"/js":      "./views/v3/js",
		"/fonts":   "./views/v3/fonts",
		"/attachs": "./attachs",
	}))

	http.Handle("/", router)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
