package stash

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RepoService interface {
	GetRepos(string) (*PagedRepos, error)
}

type repoService struct {
	*Client
}

func (r *repoService) GetRepos(projectKey string) (*PagedRepos, error) {
	req, err := r.createReq(
		"GET", fmt.Sprintf("/rest/api/1.0/projects/%s/repos", projectKey), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var b PagedRepos
	err = json.NewDecoder(resp.Body).Decode(&b)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &b, nil
}
