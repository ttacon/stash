package stash

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService interface {
	GetUsers() ([]*User, error)
	GetUser(name string) (*User, error)
	UpdateUserPassword(old, new, confirm string) error
}

type userService struct {
	*Client
}

func (u *userService) GetUsers() ([]*User, error) {
	return nil, nil
}

func (u *userService) GetUser(name string) (*User, error) {
	req, err := u.createReq(
		"GET", fmt.Sprintf("/rest/api/1.0/users/%s", name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var us User
	err = json.NewDecoder(resp.Body).Decode(&us)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &us, nil
}

func (u *userService) UpdateUserPassword(old, new, confirm string) error {
	return nil
}
