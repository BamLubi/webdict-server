package entity

import "time"

type Literature struct {
	Id         string    `json:"id"`
	Eng        string    `json:"eng"`
	Zh         string    `json:"zh"`
	Rate       string    `json:"rate"`
	Num        string    `json:"num"`
	Available  string    `json:"available"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (l Literature) TableName() string {
	return "literature"
}
