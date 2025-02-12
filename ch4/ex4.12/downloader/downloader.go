// Exercise 4.12
// downloads json data
// https://xkcd.com/571/info.0.json

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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

const URL = "https://xkcd.com"

var start = flag.Int("start", 1, "Start comic number")       // стартовый номер комикса
var end = flag.Int("end", 5, "End comic number (including)") // конечный немер комикса

var comics []Comic // collect all comics here (сюда собираем все комиксы)
var comic *Comic

func main() {
	flag.Parse()

	for comicNum := *start; comicNum <= *end; comicNum++ {
		err := downloadComic(comicNum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "downloadComic: %v\n", err)
			os.Exit(1)
		}
	}

	for _, comic := range comics {
		fmt.Printf("num: %d\n", comic.Num)
		fmt.Printf("img: %s\n", comic.Img)
		fmt.Printf("title: \"%s\"\n\n", comic.Title)
	}

	err := createFolderAndGoIntoIt("comics")
	if err != nil {
		fmt.Fprintf(os.Stderr, "createFolderAndGoIntoIt: %v\n", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(comics, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Marshal: %v\n", err)
		os.Exit(1)
	}

	err = createJsonFile(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "createJsonFile: %v\n", err)
		os.Exit(1)
	}
}

// createJsonFile
// create a file and write data to it
func createJsonFile(b []byte) error {
	now := time.Now().Format("02_01_2006_15_04")

	f, err := os.Create(now + ".json")
	if err != nil {
		return fmt.Errorf("os.Create: %v\n", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("f.Write: %v\n", err)
	}
	f.Close()

	return nil
}

func createFolderAndGoIntoIt(name string) error {
	err := os.Chdir("..") // go to the parent directory (переходим на родительскую директорию)
	if err != nil {
		return fmt.Errorf("os.Chdir: %v\n", err)
	}

	err = os.Mkdir(name, 0755) // create a comics directory (создаём директорию комиксов)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("os.Mkdir: %v\n", err)
	}

	err = os.Chdir(name)
	if err != nil {
		return fmt.Errorf("os.Chdir: %v\n", err)
	}
	return nil
}

func downloadComic(comicNum int) error {
	fmt.Printf("download comic: %d...\n", comicNum)
	resp, err := http.Get(URL + "/" + strconv.Itoa(comicNum) + "/info.0.json")
	if err != nil {
		return fmt.Errorf("http.Get: %s\n", err) // maybe pages is ending
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)
		return fmt.Errorf("resp.StatusCode != http.StatusOK\n")
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)

	err = json.Unmarshal(b, &comic)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %v\n", err)
	}

	comics = append(comics, *comic)
	time.Sleep(3000 * time.Millisecond) // preventing blocking http request (предотвращаем блокировку http запроса)
	return nil
}
