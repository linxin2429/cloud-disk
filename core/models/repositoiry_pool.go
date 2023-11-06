package models

import "time"

type RepositoryPool struct {
	Id       int
	Identity string
	Hash     string
	Name     string
	Ext      string
	Size     int64
	Path     string
	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
	DeleteAt time.Time `xorm:"deleted" json:"delete_at"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
