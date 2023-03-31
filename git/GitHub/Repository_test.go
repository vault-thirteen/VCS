package api

import (
	"testing"

	ver "github.com/vault-thirteen/VCS/common/Version"
	"github.com/vault-thirteen/tester"
)

func Test_NewRepository(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var repo *Repository

	// Test #1. Positive.
	repo, err = NewRepository("vault-thirteen", "VCS")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(repo.owner, "vault-thirteen")
	aTest.MustBeEqual(repo.path, "VCS")

	// Test #2. Negative. Bad owner.
	repo, err = NewRepository("/test", "repo")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(repo, (*Repository)(nil))

	// Test #3. Negative. Bad repository.
	repo, err = NewRepository("owner", `repo\xyz`)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(repo, (*Repository)(nil))
}

func Test_ListTags(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var repo *Repository

	repo, err = NewRepository("vault-thirteen", "VCS")
	aTest.MustBeNoError(err)

	var tags []*Tag
	tags, err = repo.ListTags()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(tags) > 0, true)
}

func Test_ListVersions(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var repo *Repository

	repo, err = NewRepository("vault-thirteen", "VCS")
	aTest.MustBeNoError(err)

	var vers []*ver.Version
	vers, err = repo.ListVersions()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(vers) > 0, true)
}

func Test_ListCleanVersions(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var repo *Repository

	repo, err = NewRepository("vault-thirteen", "VCS")
	aTest.MustBeNoError(err)

	var cleanVersions []*ver.Version
	cleanVersions, err = repo.ListCleanVersions()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(cleanVersions) > 0, true)
}
