package model

import (
	"time"
)

type BaseModel struct {
	ID uint64 `json:"id,omitempty"` 
}


// Unix 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

