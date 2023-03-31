package api

import (
	"encoding/json"

	ver "github.com/vault-thirteen/VCS/common/Version"
)

type Tag struct {
	Name       string  `json:"name"`
	Commit     *Commit `json:"commit"`
	ZipBallURL string  `json:"zipball_url"`
	TarBallURL string  `json:"tarball_url"`
	NodeId     string  `json:"node_id"`

	Version    *ver.Version `json:"-"`
	HasVersion bool         `json:"-"`
}

func ParseTags(data []byte) (tags []*Tag, err error) {
	tags = make([]*Tag, 0)
	err = json.Unmarshal(data, &tags)
	if err != nil {
		return nil, err
	}

	for i, _ := range tags {
		tags[i].Version, err = ver.New(tags[i].Name)
		if err != nil {
			tags[i].HasVersion = false
		} else {
			tags[i].HasVersion = true
		}
	}

	return tags, nil
}
