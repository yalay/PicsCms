package controllers

import (
	"conf"
	"html/template"
	"models"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/go-ozzo/ozzo-routing"
)

const (
	kArticleNumPerPage = 25
	kRelateArticleNum  = 10
)

var totalArticles *TotalArticles
var totalCates *TotalCates

func init() {
	totalArticles = NewTotalArticles(conf.RootPath())
	go totalArticles.TotalSync()

	totalCates = NewTotalCates()
	go totalCates.TotalSync()
}

func ErrorHandler(c *routing.Context) error {
	http.Redirect(c.Response, c.Request, "/", http.StatusMovedPermanently)
	return nil
}

func HomeHandler(c *routing.Context) error {
	cates := totalCates.TotalQuery()
	cateArticles := make(map[int][]*models.Article, len(cates))
	for _, cate := range cates {
		cateArticles[cate.Id] = totalArticles.QueryByCate(cate.Id, 0, 5)
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "home", map[string]interface{}{
		"webName":        conf.WebName(),
		"webKeywords":    conf.WebKeywords(),
		"webDesc":        conf.WebDesc(),
		"cid":            0,
		"totalCates":     cates,
		"sliderArticles": totalArticles.QueryByCate(2, 0, 5),
		"cateArticles":   cateArticles,
	})
}

func ArticleHandler(c *routing.Context) error {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article := totalArticles.SingleQuery(articleId)
	if article == nil || len(article.Attachs) == 0 {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "article not found")
	}

	oriPid := c.Param("pid")
	if oriPid == "p" || oriPid == "n" {
		// 上一篇或者下一篇
		preId, nextId := totalArticles.ClosestArticles(article.Cid, articleId)
		switch oriPid {
		case "p":
			preUrl := conf.GenArticleUrl(preId)
			http.Redirect(c.Response, c.Request, preUrl, http.StatusFound)
		case "n":
			nextUrl := conf.GenArticleUrl(nextId)
			http.Redirect(c.Response, c.Request, nextUrl, http.StatusFound)
		}
		return nil
	}

	pageId, _ := strconv.Atoi(oriPid)
	if articleId <= 0 || pageId < 0 {
		return conf.Render.Text(c.Response,
			http.StatusBadRequest, "invalid param")
	}

	cate := totalCates.SingleQuery(article.Cid)
	if cate == nil {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "category not found")

	}

	if pageId == 0 {
		pageId = 1
	}

	articleHomeUrl := conf.GenArticleUrl(articleId)
	pathExt := path.Ext(articleHomeUrl)
	page := &Page{
		TotalNum:  len(article.Attachs),
		CurNum:    pageId,
		SizeNum:   kRelateArticleNum,
		UrlPrefix: strings.TrimSuffix(articleHomeUrl, pathExt),
		UrlSuffix: pathExt,
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "article", map[string]interface{}{
		"id":          article.Id,
		"title":       article.Title,
		"keywords":    article.Keywords,
		"attachNum":   len(article.Attachs),
		"pageId":      pageId,
		"publishTime": article.PublishTime.Format("2006-01-02 15:04"),
		"totalCates":  totalCates.TotalQuery(),
		"cName":       cate.Name,
		"cEngName":    cate.EngName,
		"cid":         cate.Id,
		"file":        article.Attachs[pageId-1],
		"preUrl":      page.PreUrl(),
		"nextUrl":     page.NextUrl(),
		"pagination":  template.HTML(page.Html()),
		"relates": func() []*models.Article {
			articleIds := totalArticles.totalTags.Relate(article.Keywords)
			if len(articleIds) == 0 {
				return nil
			}

			// 排除自己
			for i, articleId := range articleIds {
				if articleId == article.Id {
					if i >= len(articleIds)-1 {
						articleIds = articleIds[:len(articleIds)-1]
						break
					}

					articleIds = append(articleIds[0:i], articleIds[i+1:]...)
				}
			}
			return totalArticles.MultiQuery(articleIds)
		}(),
		"tags": strings.Split(article.Keywords, ","),
	})
}

func CateHandler(c *routing.Context) error {
	cateName := c.Param("cate")
	cate := totalCates.SingleQueryByName(cateName)
	if cate == nil {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "cate not found")
	}

	pageId, _ := strconv.Atoi(c.Param("pid"))
	if pageId <= 0 {
		pageId = 1
	}

	articles := totalArticles.QueryByCate(cate.Id, (pageId-1)*kArticleNumPerPage, kArticleNumPerPage)
	if len(articles) == 0 {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "article not found")
	}

	cateHomeUrl := conf.GenCateUrl(cateName)
	pathExt := path.Ext(cateHomeUrl)
	page := &Page{
		TotalNum:  totalArticles.SumByCate(cate.Id)/kArticleNumPerPage + 1,
		CurNum:    pageId,
		SizeNum:   kArticleNumPerPage,
		UrlPrefix: strings.TrimSuffix(cateHomeUrl, pathExt),
		UrlSuffix: pathExt,
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "category", map[string]interface{}{
		"cid":        cate.Id,
		"cKeywords":  cate.Keywords,
		"pageId":     pageId,
		"totalCates": totalCates.TotalQuery(),
		"cName":      cate.Name,
		"cDesc":      cate.Desc,
		"cArticles":  articles,
		"pagination": template.HTML(page.Html()),
	})

	return nil
}

func TagsHandler(c *routing.Context) error {
	tag := c.Param("tag")
	if tag == "" {
		return conf.Render.Text(c.Response,
			http.StatusBadRequest, "invalid param")
	}

	articleIds := totalArticles.totalTags.Relate(tag, kArticleNumPerPage)
	if len(articleIds) == 0 {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "article not found")
	}

	tagHomeUrl := conf.GenTagUrl(tag)
	pathExt := path.Ext(tagHomeUrl)
	page := &Page{
		TotalNum:  1,
		CurNum:    1,
		SizeNum:   kArticleNumPerPage,
		UrlPrefix: strings.TrimSuffix(tagHomeUrl, pathExt),
		UrlSuffix: pathExt,
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "tag", map[string]interface{}{
		"webName":    conf.WebName(),
		"tag":        tag,
		"cid":        99,
		"totalCates": totalCates.TotalQuery(),
		"tArticles":  totalArticles.MultiQuery(articleIds),
		"pagination": template.HTML(page.Html()),
	})
	return nil
}
