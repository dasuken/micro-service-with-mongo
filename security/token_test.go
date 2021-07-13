package security

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestNewToken(t *testing.T) {
	userId := bson.NewObjectId()
	token, err := NewToken(userId.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewTokenPayload(t *testing.T) {
	userId := bson.NewObjectId()
	token, err := NewToken(userId.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := NewTokenPayload(token)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
	assert.Equal(t, userId.Hex(), payload.UserId)
}