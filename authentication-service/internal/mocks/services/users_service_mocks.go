package mocks

import (
	"authentication/internal/services"
	"time"
)

type UsersServiceMocks struct{}

func NewUsersServiceMocks() *UsersServiceMocks {
	return &UsersServiceMocks{}
}

func (u *UsersServiceMocks) GetAll() ([]*services.User, error) {
	users := []*services.User{}

	return users, nil
}

func (u *UsersServiceMocks) GetByEmail(email string) (*services.User, error) {
	user := services.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (u *UsersServiceMocks) GetOne(id int) (*services.User, error) {
	user := services.User{
		ID:        1,
		FirstName: "First",
		LastName:  "Last",
		Email:     "me@here.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (u *UsersServiceMocks) Update(user services.User) error {
	return nil
}

func (u *UsersServiceMocks) DeleteByID(id int) error {
	return nil
}

func (u *UsersServiceMocks) Insert(user services.User) (int, error) {
	return 2, nil
}

func (u *UsersServiceMocks) ResetPassword(password string, user services.User) error {
	return nil
}

func (u *UsersServiceMocks) PasswordMatches(plainText string, user services.User) (bool, error) {
	return true, nil
}
