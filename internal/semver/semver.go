// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semver

import (
	"strconv"

	gversion "github.com/hashicorp/go-version"
)

// LessThan returns whether or not the first version string is less than the second
// according to Semantic Versioning rules (https://semver.org/).
func LessThan(s1, s2 string) bool {
	v1, v2, err := parseVersions(s1, s2)

	if err != nil {
		return false
	}

	return v1.LessThan(v2)
}

// GreaterThanOrEqual returns whether or not the first version string is greater than or equal
// to the second according to Semantic Versioning rules (https://semver.org/).
func GreaterThanOrEqual(s1, s2 string) bool {
	v1, v2, err := parseVersions(s1, s2)

	if err != nil {
		return false
	}

	return v1.GreaterThanOrEqual(v2)
}

// MajorMinor returns the major and minor version string.
func MajorMinor(s string) (string, error) {
	v, err := gversion.NewVersion(s)

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(v.Segments64()[0], 10) + `.` + strconv.FormatInt(v.Segments64()[1], 10), nil
}

func parseVersions(s1, s2 string) (*gversion.Version, *gversion.Version, error) {
	v1, err := gversion.NewVersion(s1)

	if err != nil {
		return nil, nil, err
	}

	v2, err := gversion.NewVersion(s2)

	if err != nil {
		return nil, nil, err
	}

	return v1, v2, nil
}
