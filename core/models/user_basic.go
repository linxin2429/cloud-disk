package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserBasic struct {
	Id       int       `json:"id"`
	Identity string    `json:"identity"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
	DeleteAt time.Time `xorm:"deleted" json:"delete_at"`
}

func (u UserBasic) TableName() string {
	return "user_basic"
}

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

const JwtSecret = "cloud_disk_key"
const CaptchaLength = 6
const CaptchaExpiration = 300
