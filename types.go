package stash

type Group struct {
	Name       string `json:"name"`
	Deleteable bool   `json:"deleteable,omitempty"`
}

type PagedGroup struct {
	Size       int      `json:"size"`
	Limit      int      `json:"limit"`
	Start      int      `json:"start"`
	IsLastPage bool     `json:"isLastPage"`
	Values     []*Group `json:"values"`
}

type PagedUsers struct {
	Size       int     `json:"size"`
	Limit      int     `json:"limit"`
	Start      int     `json:"start"`
	IsLastPage bool    `json:"isLastPage"`
	Values     []*User `json:"values"`
}

type User struct {
	Active                      bool                           `json:"active"`
	Slug                        string                         `json:"slug"`
	Deleteable                  bool                           `json:"deleatable"`
	ID                          int                            `json:"id"`
	DisplayName                 string                         `json:"displayName"`
	MutableDetails              bool                           `json:"mutableDetails"`
	Link                        map[string]string              `json:"link"`
	Name                        string                         `json:"name"`
	Type                        string                         `json:"type"`
	DirectoryName               string                         `json:"directoryName"`
	LastAuthenticationTimestamp int                            `json:"lastAuthenticationTimestamp"`
	Links                       map[string][]map[string]string `json:"links"`
	EmailAddress                string                         `json:"emailAddress"`
	MutableGroups               bool                           `json:"mutableGroups"`
}
