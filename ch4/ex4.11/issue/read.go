package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Read() (*Issue, error) {

	req, err := http.NewRequest(GET, issueLink(), nil)
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

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}

	var issue *Issue

	err = json.Unmarshal(bs, &issue)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return issue, nil
}
