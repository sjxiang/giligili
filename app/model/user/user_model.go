package user

import (
	"gorm.io/gorm"

	"giligili/app/model"
)

// User 用户模型
type User struct {
	gorm.Model

	model.BaseModel
	UserName       string ``
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	model.CommonTimestampsField
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12

	// Active 激活用户
	Active string = "active"
	
	// Inactive 未激活用户
	
	Inactive string = "inactive"
	
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

