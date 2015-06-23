package stash

type Group struct {
	Name       string `json:"name"`
	Deleteable bool   `json:"deleteable,omitempty"`
}
