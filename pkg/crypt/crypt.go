package crypt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"inventory/domain"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user domain.User) (string, error) {
	usr := map[string]interface{}{
		"email": user.Email,
		"id":    user.Id.String(),
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = usr
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte(viper.GetString("SIGNING_KEY")))

	if err != nil {
		return "", err
	}

	return t, nil
}
