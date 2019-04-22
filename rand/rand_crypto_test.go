package rand

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	t.Log(GenerateRandomBytes(64))
	t.Log(GenerateRandomString(64))
}
