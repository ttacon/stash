package stash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const userAgent = "GoStashClient-1.0"

type Client struct {
	Username string
	Password string
	BaseURL  string
}

func NewClient(username, password, base string) *Client {
	return &Client{
		Username: username,
		Password: password,
		BaseURL:  base,
	}
}

func (c *Client) GroupService() GroupService {
	return &groupService{Client: c}
}

type GroupService interface {
	CreateGroup(name string) (*Group, error)
	GetGroup(name string) (*Group, error)
	GetGroups(filter string) ([]*Group, error)
}

func (c *Client) createReq(method, urlStr string, body interface{}) (*http.Request, error) {
	// this method is based off
	// https://github.com/google/go-github/blob/master/github/github.go:
	// NewRequest as it's a very nice way of doing this
	_, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// This is useful as this functionality works the same for the actual
	// BASE_URL and the download url (TODO(ttacon): insert download url)
	// this seems to be failing to work not RFC3986 (url resolution)
	//resolvedUrl := c.BaseUrl.ResolveReference(parsedUrl)
	resolvedUrl, err := url.Parse(c.BaseURL + urlStr)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, resolvedUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	// TODO(ttacon): identify which headers we should add
	// e.g. "Accept", "Content-Type", "User-Agent", etc.
	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(c.Username, c.Password)
	return req, nil
}

type groupService struct {
	*Client
}

func (g *groupService) CreateGroup(name string) (*Group, error) {
	req, err := g.createReq(
		"POST", fmt.Sprintf("/rest/api/1.0/admin/groups?name=%s", name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var r Group
	err = json.NewDecoder(resp.Body).Decode(&r)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (g *groupService) GetGroups(filter string) ([]*Group, error) {
	req, err := g.createReq(
		"GET", fmt.Sprintf("/rest/api/1.0/admin/groups?filter=%s", filter), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var r PagedGroup
	err = json.NewDecoder(resp.Body).Decode(&r)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return r.Values, nil
}

func (g *groupService) GetGroup(name string) (*Group, error) {
	groups, err := g.GetGroups(name)
	if err != nil {
		return nil, err
	} else if len(groups) == 0 {
		return nil, fmt.Errorf("no group found")
	} else if len(groups) != 1 {
		// TODO(ttacon): should we try to find an exact match among the groups though?
		return nil, fmt.Errorf("more than one group found")
	}
	return groups[0], nil
}
