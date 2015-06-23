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
