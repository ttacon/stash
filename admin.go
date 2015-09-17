package stash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AdminService interface {
	// Groups
	CreateGroup(name string) (*Group, error)
	GetGroup(name string) (*Group, error)
	DeleteGroup(name string) error
}

type adminService struct {
	*Client
}

func (a *adminService) CreateGroup(name string) (*Group, error) {
	req, err := a.createReq(
		"POST", fmt.Sprintf("/rest/api/1.0/admin/groups?name=%s", name), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(req.URL)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var g Group
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	fmt.Println(string(data))
	err = json.NewDecoder(bytes.NewBuffer(data)).Decode(&g)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (a *adminService) GetGroup(name string) (*Group, error) {
	return nil, nil
}

func (a *adminService) DeleteGroup(name string) error {
	return nil
}
