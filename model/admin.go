package model

import (
	"blog/base"
	"time"
)

type LevelEnum uint

const (
	Normal        LevelEnum = 1
	Administrator LevelEnum = 2
)

// Admin 管理员
type Admin struct {
	base.Model
	Email         string    `json:"email"`
	Nickname      string    `json:"nickname"`
	Password      string    `json:"password"`
	Level         LevelEnum `json:"level"`
	LastLoginTime time.Time `json:"last_login_time"`
}
