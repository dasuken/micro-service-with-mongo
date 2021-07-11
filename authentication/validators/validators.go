package validators

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"microservices/pb"
	"strings"
)

var (
	ErrInvalidUserId = errors.New("invalid userId")
	ErrEmptyName     = errors.New("name can't be empty")
	ErrEmptyEmail    = errors.New("email can't be empty")
	ErrEmptyPassword = errors.New("password can't be empty")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrSignInFailed = errors.New("signin failed")
)

func ValidateSignUp(req *pb.User) error {
	if !bson.IsObjectIdHex(req.Id) {
		return ErrInvalidUserId
	}

	if req.Name == "" {
		return ErrEmptyName
	}

	if req.Email == "" {
		return ErrEmptyEmail
	}

	if req.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}