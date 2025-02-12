// Exercise 4.12

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Comic struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year"`
	SaveTitle  string `json:"save_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}

var search = flag.String("search", "", "Search word")
var files []string

func main() {
	flag.Parse()

	if len(*search) == 0 {
		fmt.Fprint(os.Stderr, "\"search\" flag value is empty\n")
		os.Exit(1)
	}

	err := os.Chdir("../comics")
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Chdir: %v\n", err)
		os.Exit(1)
	}

	err = filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
		if err == nil && filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})

	for _, jsonFile := range files {
		fmt.Printf("start scan file: %s\n", jsonFile)
		f, err := os.Open(jsonFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Open: %v\n", err)
			os.Exit(1)
		}

		var t = make([]byte, 1024*1024)
		n, err := f.Read(t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "f.Read: %v\n", err)
			os.Exit(1)
		}

		var comics []Comic
		err = json.Unmarshal(t[:n], &comics)
		if err != nil {
			fmt.Fprintf(os.Stderr, "json.Unmarshal: %v\n", err)
			os.Exit(1)
		}

		for _, comic := range comics {
			if strings.Contains(comic.Title, *search) {
				fmt.Printf("\nTranscript: %s\n", comic.Transcript)
				fmt.Printf("title: %s\n\n", comic.Title)
			}
		}
	}
}
