package utils

import (
	"cloud_disk/core/models"
	"math/rand"
	"time"
)

func RandCode() string {
	s := "1234567890"
	code := make([]byte, models.CaptchaLength)

	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < models.CaptchaLength; i++ {
		code[i] = s[randSeed.Intn(len(s))]
	}
	return string(code)
}
