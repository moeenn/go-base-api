package password

import (
	"testing"
)

func TestHashAndVerify(t *testing.T) {
	pwd := "Apple_123123123"
	hash, err := Hash(pwd)
	if err != nil {
		t.Error("hashing returned an error", err)
	}

	verified := VerifyPassword(pwd, hash)
	if !verified {
		t.Error("failed password verification")
	}
}
