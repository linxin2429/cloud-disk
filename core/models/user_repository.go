package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int
	RepositoryIdentity string
	Name               string
	Ext                string
	CreateAt           time.Time `xorm:"created" json:"create_at"`
	UpdateAt           time.Time `xorm:"updated" json:"update_at"`
	DeleteAt           time.Time `xorm:"deleted" json:"delete_at"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}

const DefaultPageSize = 20
