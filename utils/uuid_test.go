package utils

import "testing"

func TestUUID(t *testing.T) {
	t.Log(GenerateUUID())
	t.Log(len(GenerateUUID()))
}
