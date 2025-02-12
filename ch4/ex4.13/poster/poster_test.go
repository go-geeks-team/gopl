package main

import (
	"flag"
	"net"
	"os"
	"testing"
)

// $ go test -v -apikey=<API_KEY>
// $ go test -v -apikey=<API_KEY> -title="The Matrix"
func TestDownloadPoster(t *testing.T) {

	_, err := net.Dial("tcp", "google.com:http") // check internet connection
	if err != nil {
		t.Fatalf("net.Dial: %v\n", err)
	}

	err = flag.Set("title", "The Mask")
	if err != nil {
		t.Fatalf("flag.Set: %v\n", err)
	}

	posterFile, err := DownloadPoster()
	if err != nil {
		t.Fatalf("DownloadPoster: %v\n", err)
	}

	wd, _ := os.Getwd()
	t.Logf("wd: %s\n", wd)

	_, err = os.OpenFile(posterFile.Name(), os.O_RDONLY, 0755) // open the file read-only.
	if err != nil {
		t.Fatalf("os.OpenFile: %v\n", err)
	}
	posterFile.Close()

	t.Logf("file: %s is downloaded (or updated)\n", posterFile.Name())
}
