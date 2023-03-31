package ver

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/vault-thirteen/auxie/number"
)

const (
	ErrPartsCount              = "parts count error: %v vs %v"
	ErrVersionsAreIncomparable = "versions are incomparable"
)

type Version struct {
	Major   uint
	Minor   uint
	Patch   uint
	Postfix string
}

var allowedVersionStringPrefixes = []string{
	"ver_", "Ver_", "VER_",
	"ver.", "Ver.", "VER.",
	"ver", "Ver", "VER",
	"v.", "V.",
	"v", "V",
}

func New(versionStr string) (v *Version, err error) {
	var tmp = versionStr
	for _, prefix := range allowedVersionStringPrefixes {
		tmp, _ = strings.CutPrefix(tmp, prefix)
	}

	parts := strings.Split(tmp, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf(ErrPartsCount, 3, len(parts))
	}

	v = &Version{}

	v.Major, err = number.ParseUint(parts[0])
	if err != nil {
		return nil, err
	}

	v.Minor, err = number.ParseUint(parts[1])
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("^\\d+")
	idx := re.FindStringIndex(parts[2])

	v.Patch, err = number.ParseUint(parts[2][idx[0]:idx[1]])
	if err != nil {
		return nil, err
	}

	v.Postfix = parts[2][idx[1]:]

	return v, nil
}

// IsClean checks whether the version has only numeric information without any
// postfixes. For example, 'v0.1.2' version is clean, but 'v0.1.2-rc5' is not
// clean.
func (v *Version) IsClean() bool {
	return len(v.Postfix) == 0
}

// CleanVersions filters an array of versions leaving only clean versions.
func CleanVersions(in []*Version) (out []*Version) {
	out = make([]*Version, 0, len(in))

	for _, x := range in {
		if x.IsClean() {
			out = append(out, x)
		}
	}

	return out
}

// IsEqualTo tells if the version is equal to a specified version.
func (v *Version) IsEqualTo(that *Version) (isEqual bool) {
	return (v.Major == that.Major) &&
		(v.Minor == that.Minor) &&
		(v.Patch == that.Patch) &&
		(v.Postfix == that.Postfix)
}

// IsGreaterThan tells if the version is greater than a specified version.
func (v *Version) IsGreaterThan(that *Version) (isGreater bool, err error) {
	if v.Major > that.Major {
		return true, nil
	} else if v.Major < that.Major {
		return false, nil
	}

	// Major versions are equal. Dig deeper.
	if v.Minor > that.Minor {
		return true, nil
	} else if v.Minor < that.Minor {
		return false, nil
	}

	// Minor versions are equal. Dig deeper.
	if v.Patch > that.Patch {
		return true, nil
	} else if v.Patch < that.Patch {
		return false, nil
	}

	// Patch versions are equal.
	// We can not compare version further.
	return false, errors.New(ErrVersionsAreIncomparable)
}