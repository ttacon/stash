package stash

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ttacon/pretty"
)

type PullRequestService interface {
	GetPRs(project, repo string) ([]*PullRequest, error)
	CreatePRFromSrcRef(proj, repo, srcRef, title, desc string, reviewers []string) (*PullRequest, error)
}

type pullRequestService struct {
	*Client
}

func (p *pullRequestService) GetPRs(project, repo string) ([]*PullRequest, error) {
	return nil, nil
}

func (p *pullRequestService) CreatePRFromSrcRef(proj, repo, srcRef, title, desc string, reviewers []string) (*PullRequest, error) {

	var repoData = map[string]interface{}{
		"slug": repo,
		"name": nil,
		"project": map[string]interface{}{
			"key": proj,
		},
	}

	var reviewerData = make([]map[string]interface{}, len(reviewers))
	for i, reviewer := range reviewers {
		reviewerData[i] = map[string]interface{}{
			"user": map[string]interface{}{
				"name": reviewer,
			},
		}
	}

	var prData = map[string]interface{}{
		"title":       title,
		"description": desc,
		"state":       "OPEN",
		"open":        true,
		"closed":      false,
		"fromRef": map[string]interface{}{
			"id":         "refs/heads/" + srcRef,
			"repository": repoData,
		},
		"toRef": map[string]interface{}{
			"id":         "refs/heads/master",
			"repository": repoData,
		},
		"locked":    false,
		"reviewers": reviewerData,
	}

	pretty.Println(prData)

	req, err := p.createReq(
		"POST",
		fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/pull-requests",
			proj, repo),
		prData,
	)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var pr PullRequest
	err = json.NewDecoder(resp.Body).Decode(&pr)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

type PullRequest struct {
	ID           int            `json:"id"`
	Version      int            `json:"version"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	State        string         `json:"state"`
	Open         bool           `json:"open"`
	Closed       bool           `json:"closed"`
	CreatedDate  int            `json:"createdDate"`
	UpdatedDate  int            `json:"updatedDate"`
	FromRef      RepoRef        `json:"fromRef"`
	ToRef        RepoRef        `json:"toRef"`
	Locked       bool           `json:"locked"`
	Author       UserWithRole   `json:"author"`
	Reviewers    []UserWithRole `json:"reviewers"`
	Participants []UserWithRole `json:"participants"`
	Link         LinkType       `json:"link"`
	Links        Links          `json:"links"`
}

type UserWithRole struct {
	User     User   `json:"user"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
}

type RepoRef struct {
	ID              string     `json:"id"`
	DisplayID       string     `json:"displayID"`
	LatestChangeset string     `json:"latestChangeset"`
	LatestCommit    string     `json:"latestCommit"`
	Repository      Repository `json:"repository"`
}

type Repository struct {
	Slug          string   `json:"slug"`
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	SCMID         string   `json:"scmId"`
	State         string   `json:"state"`
	StatusMessage string   `json:"statusMessage"`
	Forkable      bool     `json:"forkable"`
	Project       Project  `json:"project"`
	Public        bool     `json:"public"`
	CloneURL      string   `json:"cloneUrl"`
	Link          LinkType `json:"link"`
	Links         Links    `json:"links"`
}

type Project struct {
	Key         string   `json:"key"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Public      bool     `json:"public"`
	Type        string   `json:"type"`
	Link        LinkType `json:"link"`
	Links       Links    `json:"links"`
}

type LinkType struct {
	URL string `json:"url"`
	Rel string `json:"rel"`
}

type Links struct {
	Self  []Href `json:"self"`
	Clone []Href `json:"clone"`
}

type Href struct {
	Href string `json:"href"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Rel  string `json:"rel"`
}
