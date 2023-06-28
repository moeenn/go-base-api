package jwt

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateAndValidate(t *testing.T) {
	key := "some_random_secret_key"
	var expMinutes uint = 15

	payload := JWTPayload{
		UserId:    uuid.New().String(),
		UserRoles: []string{"ADMIN"},
	}

	token, err := GenerateToken(key, expMinutes, payload)
	if err != nil {
		t.Errorf("failed to generate token")
	}

	fmt.Printf("token: %v\n", token)

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

	numRoles := len(decodedPayload.UserRoles)
	if numRoles != 1 {
		t.Errorf("invalid number of decoded userRoles. Expected: %d, Got: %d", 1, numRoles)
	}

	if decodedPayload.UserRoles[0] != payload.UserRoles[0] {
		t.Errorf("unexpected user role in payload. Expected: %s, Got: %s",
			payload.UserRoles,
			decodedPayload.UserRoles,
		)
	}
}
