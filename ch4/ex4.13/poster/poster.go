package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	apikey = flag.String("apikey", "", "")
	title  = flag.String("title", "", "")
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
	ImdbID string `json:"imdbID"`
}

func main() {
	_, err := DownloadPoster()
	if err != nil {
		fmt.Fprintf(os.Stderr, "DownloadPoster: %v\n", err)
		os.Exit(1)
	}
}

func link() string {
	return fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&t=%s", *apikey, *title)
}

func checkFlags() error {
	flag.Parse()

	switch {
	case len(*apikey) == 0:
		return fmt.Errorf("flag apikey is empty\n")
	case len(*title) == 0:
		return fmt.Errorf("flag title is empty\n")
	}

	return nil
}

func DownloadPoster() (*os.File, error) {
	err := checkFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "checkFlags: %v\n", err)
		os.Exit(1)
	}

	var b = make([]byte, 1024, (1024*1024)*24) // create a byte slice with the required capacity ( создаём баётовыё срез с необходимоё ёмкостью )
	var movie Movie

	resp, err := http.Get(link())
	if err != nil {
		return nil, fmt.Errorf("http.Get: %v\n", err)
	}
	defer resp.Body.Close() // preventing resource leaking ( предотвращаем утечку ресурсов)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)
		return nil, fmt.Errorf("resp.StatusCode not equal http.StatusOK\n")
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v\n", err)
	}
	defer resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)

	err = json.Unmarshal(b, &movie)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v\n", err)
	}

	resp, err = http.Get(movie.Poster) // the second get request loads the thumbnail (image) ( второй гет запрос загружает миниатюру(изображение))
	if err != nil {
		return nil, fmt.Errorf("http.Get Poster: %v\n", err)
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v\n", err)
	}

	// creating image and writing data

	err = os.Mkdir("img", 0755)                     // creating folder for collecting images(posters)
	if err != nil && !errors.Is(err, os.ErrExist) { // avoiding file exist error (избегаем ошибку существования файла)
		return nil, fmt.Errorf("os.Mkdir: %v\n", err)
	}

	err = os.Chdir("./img") // go to the img directory ( переходим в директорию img )
	if err != nil {
		return nil, fmt.Errorf("os.Chdir: %v\n", err)
	}

	f, err := os.Create(movie.ImdbID + ".png") // create a blank image file ( содаём пустое изображение )
	if err != nil {
		return nil, fmt.Errorf("os.Create: %v\n", err)
	}

	_, err = f.Write(b) // adding data to a empty image file ( заполнняем пустой файл бинарными данными изображения )
	if err != nil {
		return nil, fmt.Errorf("f.Write: %v\n", err)
	}
	defer f.Close()

	return f, nil
}
