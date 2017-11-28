package controllers

import (
	"conf"
	"fmt"
	"html/template"
	"models"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"util"
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

func ErrorHandler(c *routing.Context, err error) error {
	return conf.Render.HTML(c.Response, http.StatusOK, "error", map[string]interface{}{
		"webName":    conf.WebName(),
		"error":      err.Error(),
		"cid":        100,
		"totalCates": totalCates.TotalQuery(),
		"tArticles":  totalArticles.QueryByCate(2, 5, 10),
	})
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
	ip:= util.RealIp(c.Request)
	ua := c.Request.UserAgent()
	if NeedAd(ip, ua) {
		return AdHandler(c)
	}

	articleId, _ := strconv.Atoi(c.Param("id"))
	article := totalArticles.SingleQuery(articleId)
	if article == nil || len(article.Attachs) == 0 {
		return fmt.Errorf("article not found:%d", articleId)
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
		return fmt.Errorf("invalid param: articleId:%d,pageId:%d", articleId, pageId)
	}

	cate := totalCates.SingleQuery(article.Cid)
	if cate == nil {
		return fmt.Errorf("category not found: cid:%d", article.Cid)
	}

	if pageId == 0 {
		pageId = 1
	}

	totalNum := len(article.Attachs)
	if pageId > totalNum {
		return fmt.Errorf("article page id not found: %d", pageId)
	}

	articleHomeUrl := conf.GenArticleUrl(articleId)
	pathExt := path.Ext(articleHomeUrl)
	page := &Page{
		TotalNum:  totalNum,
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
		"publishTime": article.PublishTime.Format("2006-01-02"),
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
		return fmt.Errorf("cate not found: %s", cateName)
	}

	pageId, _ := strconv.Atoi(c.Param("pid"))
	if pageId <= 0 {
		pageId = 1
	}

	totalNum := totalArticles.SumByCate(cate.Id)/kArticleNumPerPage + 1
	if pageId > totalNum {
		return fmt.Errorf("category page id not found: %d", pageId)
	}

	articles := totalArticles.QueryByCate(cate.Id, (pageId-1)*kArticleNumPerPage, kArticleNumPerPage)
	if len(articles) == 0 {
		return fmt.Errorf("article not found")
	}

	cateHomeUrl := conf.GenCateUrl(cateName)
	pathExt := path.Ext(cateHomeUrl)
	page := &Page{
		TotalNum:  totalNum,
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
}

func TagsHandler(c *routing.Context) error {
	tag := c.Param("tag")
	if tag == "" {
		return fmt.Errorf("empty param")
	}

	articleIds := totalArticles.totalTags.Relate(tag, kArticleNumPerPage)
	if len(articleIds) == 0 {
		return fmt.Errorf("empty article by tag: %s", tag)
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
}

func SiteMapHandler(c *routing.Context) error {
	sm := stm.NewSitemap()
	sm.Create()
	sm.SetDefaultHost(conf.WebUrl())

	sm.Add(stm.URL{"loc": "/", "changefreq": "daily"})
	cates := totalCates.TotalQuery()
	for _, cate := range cates {
		sm.Add(stm.URL{
			"loc":        conf.GenCateUrl(cate.EngName),
			"changefreq": "daily",
			"news": stm.URL{
				"publication": stm.URL{
					"name":     conf.WebName(),
					"language": "cn",
				},
				"title":    cate.Name,
				"keywords": cate.Keywords,
			},
		})
	}

	for _, cate := range cates {
		articles := totalArticles.QueryByCate(cate.Id, 0, 100)
		for _, article := range articles {
			sm.Add(stm.URL{
				"loc": conf.GenArticleUrl(article.Id),
				"news": stm.URL{
					"publication": stm.URL{
						"name":     conf.WebName(),
						"language": "cn",
					},
					"title":            article.Title,
					"keywords":         article.Keywords,
					"publication_date": article.PublishTime.Format("2006-01-02"),
				},
			})
		}
	}

	c.Response.Write(sm.XMLContent())
	return nil
}

func AdHandler(c *routing.Context) error {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article := totalArticles.SingleQuery(articleId)
	if article == nil {
		return fmt.Errorf("article not found:%d", articleId)
	}

	cate := totalCates.SingleQuery(article.Cid)
	if cate == nil {
		return fmt.Errorf("category not found: cid:%d", article.Cid)
	}

	oriUrl := c.Request.URL.String()
	return conf.Render.HTML(c.Response, http.StatusOK, "ad", map[string]interface{}{
		"title":      article.Title,
		"keywords":   article.Keywords,
		"nextUrl":    oriUrl,
		"cName":      cate.Name,
		"cEngName":   cate.EngName,
		"cid":        99,
		"totalCates": totalCates.TotalQuery(),
	})
}
