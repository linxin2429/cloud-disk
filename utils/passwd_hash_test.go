package utils

import "testing"


func TestHashPassword(t *testing.T) {
	t.Log(HashPassword("123456"))
	t.Log(HashPassword("654321"))
}
