package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
)

func TestDownloadComics(t *testing.T) {
	for comicNum := *start; comicNum <= *end; comicNum++ {
		err := downloadComic(comicNum)
		if err != nil {
			t.Fatalf("downloadComic: %v\n", err)
		}
	}
}

func TestCreateComicsFolder(t *testing.T) {
	tail := strconv.Itoa(rand.Int())
	err := createFolderAndGoIntoIt("comics" + tail)
	if err != nil {
		t.Fatalf("createFolderAndGoIntoIt: %v\n", err)
	}
}

func TestMarshalAndCreateJsonFile(t *testing.T) {
	b, err := json.MarshalIndent(comics, "", " ")
	if err != nil {
		t.Fatalf("json.Marshal: %v\n", err)
	}

	err = createJsonFile(b)
	if err != nil {
		t.Fatalf("createJsonFile: %v\n", err)
	}
}
