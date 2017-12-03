package controllers

import (
	"conf"
	"io/ioutil"
	"models"
	"path/filepath"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

type TotalArticles struct {
	rootPath        string
	articles        map[int]*models.Article
	sortedIdsByCate map[int][]int
	totalTags       *TotalTags
}

func NewTotalArticles(rootPath string) *TotalArticles {
	return &TotalArticles{
		rootPath:        rootPath,
		articles:        make(map[int]*models.Article, 0),
		sortedIdsByCate: make(map[int][]int, 0),
		totalTags:       NewTotalTags(),
	}
}

func (t *TotalArticles) TotalSync() {
	dayDirs, err := ioutil.ReadDir(t.rootPath)
	if err != nil {
		conf.Log.Error(err.Error())
		return
	}

	newArticles := make(map[int]*models.Article, 0)
	newSortedIds := make(map[int][]int, 0)
	for _, dayDir := range dayDirs {
		if !dayDir.IsDir() {
			conf.Log.Warning("not dir:" + dayDir.Name())
			continue
		}

		curDayDirPath := filepath.Join(t.rootPath, dayDir.Name())
		articleDirs, err := ioutil.ReadDir(curDayDirPath)
		if err != nil {
			conf.Log.Error(err.Error())
			continue
		}

		for _, articleDir := range articleDirs {
			if !articleDir.IsDir() {
				conf.Log.Warning("not dir:" + articleDir.Name())
				continue
			}

			curArticleDirPath := filepath.Join(curDayDirPath, articleDir.Name())
			attachs, err := ioutil.ReadDir(curArticleDirPath)
			if err != nil {
				conf.Log.Error(err.Error())
				continue
			}

			if len(attachs) == 0 {
				conf.Log.Warning("empty file:" + curArticleDirPath)
				continue
			}

			article, err := readArticleConfig(filepath.Join(curArticleDirPath, conf.ArticleProfileName()))
			if err != nil {
				conf.Log.Warning(err.Error())
				continue
			}

			conf.Log.Debug("found article:%s", article.Title)
			article.Cover = filepath.Join(curArticleDirPath, article.Cover)
			if article.HCover != "" {
				article.HCover = filepath.Join(curArticleDirPath, article.HCover)
			}

			article.Attachs = make([]string, 0, len(attachs)-1)
			for _, attach := range attachs {
				if attach.IsDir() {
					conf.Log.Warning("is dir:" + attach.Name())
					continue
				}

				if !strings.HasSuffix(attach.Name(), ".jpg") &&
					!strings.HasSuffix(attach.Name(), ".png") {
					conf.Log.Warning("not pic:" + attach.Name())
					continue
				}

				article.Attachs = append(article.Attachs, filepath.Join(curArticleDirPath, attach.Name()))
			}

			newArticles[article.Id] = article
			if curSortedIds, ok := newSortedIds[article.Cid]; !ok {
				newSortedIds[article.Cid] = []int{article.Id}
			} else {
				newSortedIds[article.Cid] = append(curSortedIds, article.Id)
			}

			t.totalTags.Insert(article.Id, article.Keywords)
		}
	}

	if len(newSortedIds) != 0 {
		for _, ids := range newSortedIds {
			sort.Sort(sort.Reverse(sort.IntSlice(ids)))
		}
	}

	t.articles = newArticles
	t.sortedIdsByCate = newSortedIds
}

func (t *TotalArticles) SingleQuery(articleId int) *models.Article {
	if len(t.articles) == 0 {
		return nil
	}

	if article, ok := t.articles[articleId]; ok {
		return article
	}
	return nil
}

func (t *TotalArticles) MultiQuery(articleIds []int) []*models.Article {
	if len(t.articles) == 0 || len(articleIds) == 0 {
		return nil
	}

	var retArticles = make([]*models.Article, 0, len(articleIds))
	for _, articleId := range articleIds {
		article := t.SingleQuery(articleId)
		if article != nil {
			retArticles = append(retArticles, article)
		}
	}

	return retArticles
}

func (t *TotalArticles) QueryByCate(cateId, startNum, count int) []*models.Article {
	if len(t.articles) == 0 || count <= 0 {
		return nil
	}

	if articleIds, ok := t.sortedIdsByCate[cateId]; !ok {
		return nil
	} else {
		totalNum := len(articleIds)
		if totalNum == 0 || startNum > totalNum {
			return nil
		}

		articles := make([]*models.Article, 0, count)
		for i := 0; i < count; i++ {
			if startNum+i >= totalNum {
				break
			}

			articleId := articleIds[startNum+i]
			if article, ok := t.articles[articleId]; ok {
				articles = append(articles, article)
			}
		}
		return articles
	}
}

func (t *TotalArticles) QueryHCoverArticles(cateId, startNum, count int) []*models.Article {
	if len(t.articles) == 0 || count <= 0 {
		return nil
	}

	if articleIds, ok := t.sortedIdsByCate[cateId]; !ok {
		return nil
	} else {
		totalNum := len(articleIds)
		if totalNum == 0 || startNum > totalNum {
			return nil
		}

		articles := make([]*models.Article, 0, count)
		for i := startNum; i < totalNum; i++ {
			articleId := articleIds[i]
			if article, ok := t.articles[articleId]; ok {
				if article.HCover != "" {
					articles = append(articles, article)
				}
			}

			if len(articles) >= count {
				break
			}
		}
		return articles
	}
}

func (t *TotalArticles) ClosestArticles(cateId, articleId int) (preId, nextId int) {
	sortedIds, ok := t.sortedIdsByCate[cateId]
	if !ok {
		return -1, -1
	}
	if len(sortedIds) == 0 || len(sortedIds) == 1 {
		return -1, -1
	}

	curIndex := reverseBinarySearch(sortedIds, articleId)
	if curIndex == 0 {
		return -1, sortedIds[1]
	}

	if curIndex > 0 && curIndex < len(sortedIds)-1 {
		return sortedIds[curIndex-1], sortedIds[curIndex+1]
	} else {
		return sortedIds[curIndex-1], -1
	}
}

func (t *TotalArticles) SumByCate(cateId int) int {
	if articles, ok := t.sortedIdsByCate[cateId]; ok {
		return len(articles)
	}
	return 0
}

func readArticleConfig(cfgFile string) (*models.Article, error) {
	article := &models.Article{}
	_, err := toml.DecodeFile(cfgFile, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func reverseBinarySearch(values []int, value int) int {
	start_index := 0
	end_index := len(values) - 1
	for start_index <= end_index {
		median := (start_index + end_index) / 2
		if values[median] > value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
	}

	if start_index == len(values) || values[start_index] != value {
		return -1
	} else {
		return start_index
	}
}
