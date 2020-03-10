package model

import (
	"sync"
	"time"
)

/// --------------------------------------------------------------------------------------------- ///
/// @Author [ETongfu].
/// @Des [Database base model].
/// --------------------------------------------------------------------------------------------- ///

type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json: "-"`
	CreatedAt time.Time  `gorm:"column:created_at" json: "-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json: "-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}

type UserInfo struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// UserList Lock and IdMap 是为了并发处理中，更新同一个变量为了保证数据的唯一性。
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// Token for user
type Token struct {
	Token string `json:"token"`
}
