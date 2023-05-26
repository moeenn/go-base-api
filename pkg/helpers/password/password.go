package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type DecodedPasswordComponents struct {
	Params PasswordParams
	Salt   []byte
	Hash   []byte
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

var params = PasswordParams{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func Hash(password string) (string, error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateSaltBytes(params.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.iterations,
		params.memory,
		params.parallelism,
		params.keyLength,
	)

	encodedHash := encodeHashToString(salt, hash)
	return encodedHash, nil
}

func encodeHashToString(salt []byte, hash []byte) string {
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		params.memory,
		params.iterations,
		params.parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash
}

func generateSaltBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func VerifyPassword(password, encodedHash string) bool {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	components, err := decodeHash(encodedHash)
	if err != nil {
		return false
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey(
		[]byte(password),
		components.Salt,
		components.Params.iterations,
		components.Params.memory,
		components.Params.parallelism,
		components.Params.keyLength,
	)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(components.Hash, otherHash) == 1 {
		return true
	}

	return false
}

func decodeHash(encodedHash string) (DecodedPasswordComponents, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return DecodedPasswordComponents{}, ErrInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return DecodedPasswordComponents{}, err
	}
	if version != argon2.Version {
		return DecodedPasswordComponents{}, ErrIncompatibleVersion
	}

	result := DecodedPasswordComponents{}
	_, err = fmt.Sscanf(
		vals[3],
		"m=%d,t=%d,p=%d",
		&result.Params.memory,
		&result.Params.iterations,
		&result.Params.parallelism,
	)

	if err != nil {
		return DecodedPasswordComponents{}, err
	}

	result.Salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return DecodedPasswordComponents{}, err
	}
	result.Params.saltLength = uint32(len(result.Salt))

	result.Hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return DecodedPasswordComponents{}, err
	}

	result.Params.keyLength = uint32(len(result.Hash))
	return result, nil
}
