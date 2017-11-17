package controllers

import (
	"sort"
	"strings"
	"sync"
)

const (
	kTopMaxNum = 10
)

type TotalTags struct {
	sync.RWMutex
	tagArticleIds map[string][]int
}

func NewTotalTags() *TotalTags {
	return &TotalTags{
		tagArticleIds: make(map[string][]int, 0),
	}
}

func (t *TotalTags) Insert(articleId int, keywords string) {
	if keywords == "" {
		return
	}

	curKeys := strings.Split(keywords, ",")
	t.Lock()
	for _, curKey := range curKeys {
		if curArticleIds, ok := t.tagArticleIds[curKey]; ok {
			curArticleIds = append(curArticleIds, articleId)
		} else {
			t.tagArticleIds[curKey] = []int{articleId}
		}
	}
	t.Unlock()
}

func (t *TotalTags) Relate(keywords string, args ...int) []int {
	if keywords == "" {
		return nil
	}

	idsFreq := NewIdsFreq()
	curKeys := strings.Split(keywords, ",")
	t.RLock()
	for _, curKey := range curKeys {
		curArticleIds, ok := t.tagArticleIds[curKey]
		if !ok || len(curArticleIds) == 0 {
			continue
		}
		idsFreq.Append(curArticleIds)
	}
	t.RUnlock()

	sort.Sort(idsFreq)
	if len(args) == 0 {
		return idsFreq.Top(kTopMaxNum)
	}

	return idsFreq.Top(args[0])
}
