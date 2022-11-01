package base

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Id int64 `json:"id"`
}
