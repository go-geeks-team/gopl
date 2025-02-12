package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Update() (*Issue, error) {
	flag.Parse()
	//var updated bool

	tmpFile, err := OpenDefaultFileManager(*desc)
	if err != nil {
		return nil, err
	}
	var iss *Issue

	//for !updated {

	for i := 60; i != 0; i-- {
		_ = os.Stdin.Truncate(1)
		fmt.Printf("\"You need to close the text editor within: (Вам нужно закрыть текстовый редактор в течении:): %d second(s)\n", i)
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Printf("Saving the file...\n")

	f, err := os.Open(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}

	data := make([]byte, 1024*8)
	fmt.Printf("%s\n", data)
	n, err := f.Read(data)
	if err != nil {
		return nil, fmt.Errorf("f.Read: %v", err)
	}

	issue := &Issue{
		Title:  *title,
		Repo:   *repo,
		Body:   string(data[:n]),
		Owner:  *owner,
		Labels: getLabels(),
	}

	// perform marshalling, i.e. perform the transformation of the Issue data structure into JSON (binary representation)
	b, err := json.Marshal(&issue) // выполняем маршаллинг т.е. выполняем преобразование структуры данных Issue в JSON(бинарное представление)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %s\n", err)
	}

	req, err := http.NewRequest(PATCH, updateIssueLink(), bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %s\n", err)
	}

	SetHeaders(req) // set the necessary http-headers (устанавливаем необходимые http-заголовки)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %v", err)
	}
	defer resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)

	iss = issue

	return iss, nil
}
