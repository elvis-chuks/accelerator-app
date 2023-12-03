package crypt

import (
	uuid "github.com/satori/go.uuid"
	"inventory/domain"
	"testing"
)

func TestHashPasswordAndCheckHashPassword(t *testing.T) {
	password, err := HashPassword("12345")

	if err != nil {
		t.Error(err)
		return
	}

	if !CheckPasswordHash("12345", password) {
		t.Error("could not confirm password hash")
	}
}

func TestGenerateToken(t *testing.T) {

	userId, err := uuid.FromString("151afdd8-888b-4d5d-9b8c-3e6ae67cedc8")

	if err != nil {
		t.Error(err)
		return
	}

	token, err := GenerateToken(domain.User{
		Email: "elvis.aganoke@gmail.com",
		Id: domain.UUID{
			UUID: userId,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if len(token) == 0 {
		t.Error("token cannot be empty")
		return
	}
}
