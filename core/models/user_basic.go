package models

import "github.com/golang-jwt/jwt/v4"

type UserBasic struct {
	Id       int    `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
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

var JwtSecret = "cloud_disk_key"