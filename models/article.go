package models

import "time"

type Article struct {
	Id          int
	Cid         int
	Title       string
	Desc        string
	Keywords    string
	Cover       string
	Attachs     []string
	PublishTime time.Time
}
