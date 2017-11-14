package controllers

import (
	"models"
	"conf"
	"sort"

	"github.com/BurntSushi/toml"
)

type TotalCates struct {
	cfgFile string
	cates map[int]*models.Category
	sortedIds []int
}

type CateConfig struct {
	Cates []*models.Category
}

func NewTotalCates(cfgFile string) *TotalCates{
	return &TotalCates{
		cfgFile:cfgFile,
		cates:make(map[int]*models.Category, 0),
		sortedIds:make([]int, 0),
	}
}

func(t *TotalCates) TotalSync() {
	cates , err := readCateConfig(t.cfgFile)
	if err != nil {
		conf.Log.Error(err.Error())
		return
	}

	for _, cate := range cates {
		t.cates[cate.Id] = cate
		t.sortedIds = append(t.sortedIds, cate.Id)
	}

	sort.Ints(t.sortedIds)
	conf.Log.Debug("TotalCates:%+v", t)
}

func (t *TotalCates) SingleQuery(cateId int) *models.Category{
	if len(t.cates) == 0 {
		return nil
	}

	if cate, ok := t.cates[cateId]; ok {
		return cate
	}

	return nil
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

func readCateConfig(cfgFile string) ([]*models.Category, error){
	cateCfg := &CateConfig{}
	_, err := toml.DecodeFile(cfgFile, cateCfg)
	if err != nil {
		return nil, err
	}

	return cateCfg.Cates, nil
}
