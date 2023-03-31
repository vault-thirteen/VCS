package ver

import (
	"fmt"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		VersionString   string
		ExpectedVersion *Version
		IsErrorExpected bool
	}

	tests := []TestData{
		{
			VersionString:   "no numbers",
			ExpectedVersion: (*Version)(nil),
			IsErrorExpected: true,
		},
		{
			VersionString: "v123.456.789quake",
			ExpectedVersion: &Version{
				Major:   123,
				Minor:   456,
				Patch:   789,
				Postfix: "quake",
			},
			IsErrorExpected: false,
		},
		{
			VersionString: "VER_0.1.2-alpha",
			ExpectedVersion: &Version{
				Major:   0,
				Minor:   1,
				Patch:   2,
				Postfix: "-alpha",
			},
			IsErrorExpected: false,
		},
		{
			VersionString: "V.1.2.3-rc4",
			ExpectedVersion: &Version{
				Major:   1,
				Minor:   2,
				Patch:   3,
				Postfix: "-rc4",
			},
			IsErrorExpected: false,
		},
	}

	type Result struct {
		Version *Version
		Error   error
	}

	var result Result

	n := 1
	for _, test := range tests {
		fmt.Print("[", n, "] ")

		result.Version, result.Error = New(test.VersionString)

		if test.IsErrorExpected {
			aTest.MustBeAnError(result.Error)
		} else {
			aTest.MustBeNoError(result.Error)
		}

		aTest.MustBeEqual(result.Version, test.ExpectedVersion)

		n++
	}
	fmt.Println()
}

func Test_IsClean(t *testing.T) {
	aTest := tester.New(t)
	var v *Version

	v = &Version{Postfix: "-alpha"}
	aTest.MustBeEqual(v.IsClean(), false)

	v = &Version{Postfix: ""}
	aTest.MustBeEqual(v.IsClean(), true)
}

func Test_CleanVersions(t *testing.T) {
	aTest := tester.New(t)

	vers := []*Version{
		{Major: 1, Minor: 2, Patch: 3, Postfix: ""},
		{Major: 4, Minor: 5, Patch: 6, Postfix: "-rc"},
		{Major: 7, Minor: 8, Patch: 9, Postfix: ""},
	}
	expectedCleanVers := []*Version{
		{Major: 1, Minor: 2, Patch: 3, Postfix: ""},
		{Major: 7, Minor: 8, Patch: 9, Postfix: ""},
	}
	aTest.MustBeEqual(CleanVersions(vers), expectedCleanVers)
}

func Test_IsEqualTo(t *testing.T) {
	aTest := tester.New(t)

	vA := &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB := &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	aTest.MustBeEqual(vA.IsEqualTo(vB), true)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 0, Minor: 2, Patch: 3, Postfix: "test"}
	aTest.MustBeEqual(vA.IsEqualTo(vB), false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 0, Patch: 3, Postfix: "test"}
	aTest.MustBeEqual(vA.IsEqualTo(vB), false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 2, Patch: 0, Postfix: "test"}
	aTest.MustBeEqual(vA.IsEqualTo(vB), false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "rc1"}
	aTest.MustBeEqual(vA.IsEqualTo(vB), false)
}

func Test_IsGreaterThan(t *testing.T) {
	aTest := tester.New(t)
	var isGreater bool
	var err error

	vA := &Version{Major: 99, Minor: 2, Patch: 3, Postfix: "test"}
	vB := &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, true)

	vA = &Version{Major: 1, Minor: 99, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, true)

	vA = &Version{Major: 1, Minor: 2, Patch: 99, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, true)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 99, Minor: 2, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 99, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	vB = &Version{Major: 1, Minor: 2, Patch: 99, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(isGreater, false)

	vA = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "999"}
	vB = &Version{Major: 1, Minor: 2, Patch: 3, Postfix: "test"}
	isGreater, err = vA.IsGreaterThan(vB)
	aTest.MustBeAnError(err)
}
