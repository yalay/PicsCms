package controllers

import (
	"conf"
	"html/template"
	"models"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-ozzo/ozzo-routing"
)

const (
	kArticleNumPerPage = 24
)

var totalArticles *TotalArticles
var totalCates *TotalCates

func init() {
	totalArticles = NewTotalArticles(conf.RootPath())
	go totalArticles.TotalSync()

	totalCates = NewTotalCates(conf.CateProfilePath())
	go totalCates.TotalSync()
}

func ArticleHandler(c *routing.Context) error {
	articleId, _ := strconv.Atoi(c.Param("id"))
	pageId, _ := strconv.Atoi(c.Param("pid"))
	if articleId <= 0 || pageId < 0 {
		return conf.Render.Text(c.Response,
			http.StatusBadRequest, "invalid param")
	}

	article := totalArticles.SingleQuery(articleId)
	if article == nil || len(article.Attachs) == 0 {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "article not found")
	}

	if pageId > len(article.Attachs) {
		// next article
		return nil
	}

	cate := totalCates.SingleQuery(article.Cid)
	if cate == nil {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "category not found")

	}

	if pageId == 0 {
		pageId = 1
	}

	attachNum := len(article.Attachs)
	page := &Page{
		ArticleId: articleId,
		TotalNum:  attachNum,
		CurNum:    pageId,
		SizeNum:   10,
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "article", map[string]interface{}{
		"id":          article.Id,
		"title":       article.Title,
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
	cateTitle := c.Param("cate")
	cate := totalCates.SingleQueryByTitle(cateTitle)
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

	return conf.Render.HTML(c.Response, http.StatusOK, "category", map[string]interface{}{
		"cid":        cate.Id,
		"pageId":     pageId,
		"totalCates": totalCates.TotalQuery(),
		"cName":      cate.Name,
		"cDesc":      cate.Desc,
		"cArticles":  articles,
		//"pagination": template.HTML(page.Html()),
	})

	return nil
}

func TagsHandler(c *routing.Context) error {
	tag := c.Param("tag")
	if tag == "" {
		return conf.Render.Text(c.Response,
			http.StatusBadRequest, "invalid param")
	}

	articleIds := totalArticles.totalTags.Relate(tag)
	if len(articleIds) == 0 {
		return conf.Render.Text(c.Response,
			http.StatusNotFound, "article not found")
	}

	return conf.Render.HTML(c.Response, http.StatusOK, "tag", map[string]interface{}{
		"tag":       tag,
		"cid":       99,
		"tArticles": totalArticles.MultiQuery(articleIds),
		//"pagination": template.HTML(page.Html()),
	})
	return nil
}
