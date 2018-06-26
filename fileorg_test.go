package main

import (
	"os"
	"reflect"
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
	info, _ := os.Stat("fileorg_test.go")
	infoArr := []os.FileInfo{info}
	filetypes := OrganizeByFiletype(infoArr)
	f := map[string][]string{".go": []string{"fileorg_test.go"}}
	eq := reflect.DeepEqual(filetypes, f)

	if eq == false {
		t.Errorf("TestOrganizeByFiletype did not yield the correct filetype. got: %#v, want: %#v", filetypes, f)
	}
}

func TestCreateSubdirectories(t *testing.T) {
	// TODO
}

func TestMain(t *testing.T) {
	// Is this necessary?
}
