package api

import (
	"errors"
	"strings"

	ver "github.com/vault-thirteen/VCS/common/Version"
)

// Repository object.
type Repository struct {
	api *API

	// owner is an owner of a repository.
	// If the repository is 'https://github.com/vault-thirteen/VCS',
	// then owner is 'vault-thirteen'.
	owner string

	// path is a path of a repository.
	// If the repository is 'https://github.com/vault-thirteen/VCS',
	// then path is 'VCS'.
	path string

	tags []Tag
}

const (
	ForbiddenChars = `/\`
)

const (
	ErrOwner = "owner's name contains forbidden symbols"
	ErrPath  = "repository's path contains forbidden symbols"
)

// NewRepository creates an object of Repository type.
// Supported APIs can be found in the GitHub's file.
func NewRepository(owner string, repository string) (r *Repository, err error) {
	if strings.ContainsAny(owner, ForbiddenChars) {
		return nil, errors.New(ErrOwner)
	}
	if strings.ContainsAny(repository, ForbiddenChars) {
		return nil, errors.New(ErrPath)
	}

	return &Repository{
		api:   GitHubAPI(),
		owner: owner,
		path:  repository,
		tags:  nil,
	}, nil
}

func (r *Repository) ListTags() (tags []*Tag, err error) {
	return r.getTags()
}

func (r *Repository) ListVersions() (versions []*ver.Version, err error) {
	return r.getVersions(false)
}

func (r *Repository) ListCleanVersions() (versions []*ver.Version, err error) {
	return r.getVersions(true)
}

func (r *Repository) getTags() (tags []*Tag, err error) {
	var jsonData []byte
	jsonData, err = r.getTagsJSON()
	if err != nil {
		return nil, err
	}

	return ParseTags(jsonData)
}

func (r *Repository) getVersions(onlyCleanVersions bool) (versions []*ver.Version, err error) {
	var tags []*Tag
	tags, err = r.getTags()
	if err != nil {
		return nil, err
	}

	versions = make([]*ver.Version, 0, len(tags))
	for _, tag := range tags {
		if tag.HasVersion {
			versions = append(versions, tag.Version)
		}
	}

	if onlyCleanVersions {
		versions = ver.CleanVersions(versions)
	}

	return versions, nil
}
