package models

import "time"

type Article struct {
	Id          int
	Cid         int
	Title       string
	Desc        string
	Keywords    string
	Cover       string
	HCover      string // 横版封面
	Attachs     []string
	PublishTime time.Time
}
