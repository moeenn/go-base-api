package jwt

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateAndValidate(t *testing.T) {
	key := "some_random_secret_key"
	var expMinutes uint = 15

	payload := JWTPayload{
		UserId:   uuid.New().String(),
		UserRole: "ADMIN",
	}

	token, err := GenerateToken(key, expMinutes, payload)
	if err != nil {
		t.Errorf("failed to generate token")
	}

	decodedPayload, err := ValidateToken(key, token.Token)
	if err != nil {
		t.Errorf("failed to decode token")
	}

	if decodedPayload.UserId != payload.UserId {
		t.Errorf("unexpected user id in payload. Expected: %s, Got: %s",
			payload.UserId,
			decodedPayload.UserId,
		)
	}

	if decodedPayload.UserRole != payload.UserRole {
		t.Errorf("unexpected user role in payload. Expected: %s, Got: %s",
			payload.UserRole,
			decodedPayload.UserRole,
		)
	}
}
