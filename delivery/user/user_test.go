package user

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"inventory/domain"
	"inventory/pkg/env"
	"inventory/pkg/logger"
	"inventory/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initenv() {
	env.LoadConfig("../../")
}

func TestHandler_Signup(t *testing.T) {

	initenv()
	l, _ := logger.Init()

	type userType struct {
		User       domain.User
		ShouldPass bool
		Reason     string
	}

	cases := []userType{
		{
			User: domain.User{
				Fullname: "Elvis Chuks",
				Email:    "elvis.aganoke@gmail.com",
				Password: "1234",
			},
			ShouldPass: false,
			Reason:     "user already exists",
		},
		{
			User: domain.User{
				Fullname: "Elvis Chuks",
				Email:    "elvis.1aganoke@gmail.com",
				Password: "1234",
			},
			ShouldPass: true,
		},
	}

	for _, i := range cases {
		jsonStr, err := json.Marshal(i.User)

		if err != nil {
			t.Error(err)
			return
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(jsonStr))

		repo := repository.SetupDb(l)

		app := fiber.New()

		router := app.Group("/api/v1/auth")

		New(router, repo.UserRepo)

		resp, err := app.Test(req, 10000)

		if err != nil {
			t.Error(err)
			return
		}

		if i.ShouldPass {
			assert.Equal(t, 200, resp.StatusCode, "should pass")
		} else {
			assert.NotEqual(t, 400, resp.StatusCode, i.Reason)
		}

		resp.Body.Close()
	}
}

func TestHandler_Signin(t *testing.T) {

	initenv()
	l, _ := logger.Init()

	type userType struct {
		User       domain.User
		ShouldPass bool
		Reason     string
	}

	cases := []userType{
		{
			User: domain.User{
				Email:    "elvis.aganoke@gmail.com",
				Password: "123elvi",
			},
			ShouldPass: false,
			Reason:     "invalid credentials",
		},
		{
			User: domain.User{
				Email:    "elvis.aganoke+1@gmail.com",
				Password: "123elvis",
			},
			ShouldPass: true,
		},
	}

	for _, i := range cases {
		jsonStr, err := json.Marshal(i.User)

		if err != nil {
			t.Error(err)
			return
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signin", bytes.NewBuffer(jsonStr))

		repo := repository.SetupDb(l)

		app := fiber.New()

		router := app.Group("/api/v1/auth")

		New(router, repo.UserRepo)

		resp, err := app.Test(req, 10000)

		if err != nil {
			t.Error(err)
			return
		}

		//var response map[string]interface{}
		//
		//err = json.NewDecoder(resp.Body).Decode(&response)
		//
		//if err != nil {
		//	t.Error(err)
		//	return
		//}
		//
		//fmt.Println(response)

		//fmt.Println(resp.StatusCode)

		if i.ShouldPass {
			assert.Equal(t, 200, resp.StatusCode, "should pass")
		} else {
			assert.Equal(t, 400, resp.StatusCode, i.Reason)
		}

		resp.Body.Close()
	}
}
