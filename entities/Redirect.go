package entities

import "encoding/json"

type Redirect struct {
	Id            string
	RedirectToURL string
}

func (r Redirect) ToHash() (map[string]any, error) {
	var hash map[string]any
	var err error

	data, err := json.Marshal(r)
	if err == nil {
		err = json.Unmarshal(data, &hash)
	}

	return hash, err
}
