package main

import (
	"testing"
)

func TestOpenDirectory(t *testing.T) {
	files := OpenDirectory()
	length := len(files)
	if length < 1 {
		t.Errorf("OpenDirectory did not find any files, got: %d, want at least: %d.", length, 1)
	}
	// This could be more robust
}

func TestOrganizeByFiletype(t *testing.T) {
	// TODO
}

func TestCreateSubdirectories(t *testing.T) {
	// TODO
}

func TestMain(t *testing.T) {
	// Is this necessary?
}
