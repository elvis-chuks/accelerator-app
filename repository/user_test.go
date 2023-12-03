package repository

import (
	"inventory/domain"
	"inventory/pkg/logger"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	initenv()
	l, _ := logger.Init()

	repo := SetupDb(l).UserRepo

	type userType struct {
		Fullname   string
		Email      string
		Password   string
		ShouldPass bool
	}

	cases := []userType{
		{
			Fullname:   "Elvis",
			Email:      "elvis+egs@gmail.com",
			Password:   "1234",
			ShouldPass: true,
		},
		{
			Fullname:   "Elvis",
			Email:      "elvis+aganoke@gmail.com",
			Password:   "1234",
			ShouldPass: true,
		},
		{
			Fullname:   "Elvis",
			Email:      "elvis+aganoke@gmail.com",
			Password:   "1234",
			ShouldPass: false,
		},
	}

	for _, i := range cases {
		_, err := repo.Create(domain.User{
			Fullname: i.Fullname,
			Email:    i.Email,
			Password: i.Password,
		})

		if i.ShouldPass {
			if err != nil {
				t.Errorf("expected test to pass \n%v", err)
			}
		} else {
			if err == nil {
				t.Errorf("expected test to fail for %s,  \n%v", i.Email, err)
			}
		}
	}

}

func TestUserRepository_GetByEmail(t *testing.T) {
	initenv()
	l, _ := logger.Init()

	repo := SetupDb(l).UserRepo

	type userType struct {
		Email      string
		ShouldPass bool
	}

	cases := []userType{
		{
			Email:      "elvis+aganoke@gmail.com",
			ShouldPass: true,
		},
		{
			Email:      "egg@gmail.com",
			ShouldPass: false,
		},
	}

	for _, i := range cases {
		user, err := repo.GetByEmail(i.Email)

		if i.ShouldPass {
			if err == nil && user == nil {
				t.Error("expected test to pass but user is nil")
			}
		} else {
			if err != nil {
				t.Errorf("expected test to fail for %s,  \n%v", i.Email, err)
			}
		}
	}

}

func TestUserRepository_GetById(t *testing.T) {
	initenv()
	l, _ := logger.Init()

	repo := SetupDb(l).UserRepo

	type userType struct {
		Id         string
		ShouldPass bool
	}

	cases := []userType{
		{
			Id:         "932a4c03-bd03-4c78-89e6-fd92aa880c71",
			ShouldPass: true,
		},
		{
			Id:         "932a4c03-bd03-4c78-89e6-fd92aa880c73",
			ShouldPass: false,
		},
	}

	for _, i := range cases {
		user, err := repo.GetById(i.Id)

		if i.ShouldPass {
			if err == nil && user == nil {
				t.Error("expected test to pass but user is nil")
			}
		} else {
			if err != nil {
				t.Errorf("expected test to fail for %s,  \n%v", i.Id, err)
			}
		}
	}
}
