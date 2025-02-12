package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

// Close
// calling GitHub API for closing issue
//
// Example:
// $ go test -v -owner="<REPOSITORY_OWNER>"
//
//	-collab="<COLLABORATOR_USERNAME"
//	-token="<COLLABORATOR_GITHUB_TOKEN>"
//	-repo="<REPOSITORY_NAME>"
//	-close
func Close() error {
	flag.Parse()

	issue := &Issue{
		Repo:  *repo,
		Owner: *owner,
		State: "closed",
	}

	// perform marshalling, i.e. perform the transformation of the Issue data structure into JSON (binary representation)
	b, err := json.Marshal(&issue) // выполняем маршаллинг т.е. выполняем преобразование структуры данных Issue в JSON(бинарное представление)
	if err != nil {
		return fmt.Errorf("json.Marshal: %s\n", err)
	}

	req, err := http.NewRequest(PATCH, updateIssueLink(), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("http.NewRequest: %s\n", err)
	}

	SetHeaders(req) // set the necessary http-headers (устанавливаем необходимые http-заголовки)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do: %v", err)
	}
	defer resp.Body.Close() // preventing resource leakage (предотвращаем утечку ресурсов)

	return nil
}
