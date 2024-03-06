package main

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	re := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if !re.MatchString(VERSION) {
		t.Error("VERSION does not match vMAJOR.MINOR.PATCH format:", VERSION)
	}
}
