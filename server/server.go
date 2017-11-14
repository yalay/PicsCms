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
	flag.BoolVar(&debug, "d", false, "weather debug")
	flag.Parse()

	conf.InitLog("./test.log", debug)
}

func main() {
	router := routing.New()
	router.Get("/cates/<id>", controllers.CateHandler)
	router.Get(`/article-<id:\d+>.html`, controllers.ArticleHandler)
	router.Get(`/article-<id:\d+>-<pid:\d+>.html`, controllers.ArticleHandler)

	// static file
	router.Get("/*", file.Server(file.PathMap{
		"/css":  "./views/v3/img",
		"/js":   "./views/v3/js",
		"/test": "./test",
	}))

	http.Handle("/", router)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
