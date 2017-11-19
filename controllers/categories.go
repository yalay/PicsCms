package controllers

import (
	"conf"
	"models"
	"sort"
)

type TotalCates struct {
	cfgFile   string
	nameIds   map[string]int
	cates     map[int]*models.Category
	sortedIds []int
}

type CateConfig struct {
	Cates []*models.Category
}

func NewTotalCates() *TotalCates {
	return &TotalCates{
		cates:     make(map[int]*models.Category, 0),
		nameIds:   make(map[string]int, 0),
		sortedIds: make([]int, 0),
	}
}

func (t *TotalCates) TotalSync() {
	for _, cate := range conf.TotalCates() {
		t.nameIds[cate.EngName] = cate.Id
		t.cates[cate.Id] = cate
		t.sortedIds = append(t.sortedIds, cate.Id)
	}

	sort.Ints(t.sortedIds)
	//conf.Log.Debug("TotalCates:%+v", t)
}

func (t *TotalCates) SingleQuery(cateId int) *models.Category {
	if len(t.cates) == 0 {
		return nil
	}

	if cate, ok := t.cates[cateId]; ok {
		return cate
	}

	return nil
}

func (t *TotalCates) SingleQueryByName(name string) *models.Category {
	if id, ok := t.nameIds[name]; !ok {
		return nil
	} else {
		return t.SingleQuery(id)
	}
}

func (t *TotalCates) TotalQuery() []*models.Category {
	if len(t.sortedIds) == 0 {
		return nil
	}

	totalCates := make([]*models.Category, 0, len(t.sortedIds))
	for _, cateId := range t.sortedIds {
		if cate, ok := t.cates[cateId]; ok {
			totalCates = append(totalCates, cate)
		}
	}
	return totalCates
}
