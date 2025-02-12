package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	data, err := Create()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("owner: %s\n", data.Owner)
	t.Logf("title: %s\n", data.Title)
	t.Logf("desc: %s\n", data.Body)
	t.Logf("created_at: %s\n", data.CreatedAt)
	t.Logf("url: %s\n", data.HTMLURL)

}
