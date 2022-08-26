package model

import (
	"time"
)

// 靠，才发现 塞了 gorm.Model 匿名字段，就不用写这些 

type BaseModel struct {
	ID uint64 `json:"id,omitempty"` 
}


// Unix 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

