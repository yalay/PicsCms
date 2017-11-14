package controllers

import (
	"sort"
	"sync"
	"strings"
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
			curArticleIds  =append(curArticleIds, articleId)
		} else {
			t.tagArticleIds[curKey] = []int{articleId}
		}
	}
	t.Unlock()
}

func (t *TotalTags) Relate(keywords string, excludeId int) []int {
	if keywords == "" {
		return nil
	}

	idsFreq := &IdsFreq{}
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

	var count int
	topIds := idsFreq.Top(kTopMaxNum+1)
	relateIds := make([]int, 0, kTopMaxNum)
	for i, topId := range topIds {
		if topId == excludeId {
			if i < len(topIds) - 1 {
				relateIds = append(relateIds, topIds[i+1:]...)
			}
			break
		}

		count++
		if count > kTopMaxNum {
			break
		}
		relateIds = append(relateIds, topId)
	}

	return relateIds
}

