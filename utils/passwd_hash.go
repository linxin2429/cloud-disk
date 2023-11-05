package utils

import (
	"cloud_disk/core/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 14

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword)
	return cost
}

func GenerateToken(id int, identity string, name string) (string, error) {
	uc := models.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, uc).SignedString([]byte(models.JwtSecret))
	if err != nil {
		return "", NewErrWrapper(err, "utils.GenerateToken")
	}

	return token, nil
}
