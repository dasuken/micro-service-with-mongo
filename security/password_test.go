package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	pass, err := EncryptPassword("1234556")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
}

func TestVerifyPassword(t *testing.T) {
	pass := "123456"
	hashed, err := EncryptPassword(pass)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	assert.NoError(t, VerifyPassword(hashed, pass))
	assert.Error(t, VerifyPassword(hashed, "234"))
	assert.Error(t, VerifyPassword(hashed, hashed))
}