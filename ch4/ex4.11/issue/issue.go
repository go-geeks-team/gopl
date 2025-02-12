package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	PATCH = "PATCH" // Аналогично PUT, но применяется только к фрагменту ресурса.
	GET   = "GET"   // Используется для запроса содержимого указанного ресурса
	POST  = "POST"  // Применяется для передачи пользовательских данных заданному ресурсу (html-форма)
)

type Issue struct {
	Number    int       `json:"number"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title,omitempty"`
	State     string    `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:"body,omitempty"`
	Owner     string    `json:"owner"`
	Repo      string    `json:"repo"`
	Labels    []string  `json:"labels,omitempty"`
}

var (
	token  = flag.String("token", "", "Please enter user token: -token=\"xxxxxxxxxx\"") // token пользователя
	collab = flag.String("collab", "", "Please enter collaborator: -collab=\"john\"")   // логин коллаборатора
	owner  = flag.String("owner", "", "Please enter repository owner: -owner=\"John\"") // владелец репозитория: -owner="John"
)

var (
	repo   = flag.String("repo", "", "user repository: -repo=\"testRepo1\"") // Название репозитория : -repo="testRepo1"
	title  = flag.String("title", "", "title for the issues")                // "Заголовок" issue
	desc   = flag.String("desc", "", "body for the issues")                  // "Тело сообщения" issue
	labels = flag.String("labels", "", "Example: []string{\"bug\"}")         // "Метка issue", например: []string{"bug"}
)

var (
	createIssue = flag.Bool("create", false, "CreateIssue Issue") // Создание issue
	updateIssue = flag.Bool("update", false, "UpdateIssue Issue") // Обновление issue
	closeIssue  = flag.Bool("close", false, "Close Issue")        // Закрытие issue
	openIssue   = flag.Bool("open", false, "Open Issue")          // Открытие issue
	readIssue   = flag.Bool("read", false, "Read Issue")          // Чтение issue
	issueNum    = flag.Int("issue_number", 0, "Issue number")     // Номер issue
)

func main() {
	flag.Parse()

	if *createIssue || *updateIssue {
		err := checkCreateUpdateFlags()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		if *createIssue { // create issue flag
			issue, err := Create() // create new issue
			if err != nil {
				fmt.Errorf("Create: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("repo: %s\n", issue.Repo)
			fmt.Printf("title: %s\n", issue.Title)
			fmt.Printf("body: %s", issue.Body)
			fmt.Printf("created_at: %s\n", issue.CreatedAt)
			fmt.Printf("labels: %v\n", issue.Labels)
		}

		if *updateIssue { // update issue flag
			issue, err := Update() // update issue
			if err != nil {
				fmt.Errorf("Update: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("repo: %s\n", issue.Repo)
			fmt.Printf("title: %s\n", issue.Title)
			fmt.Printf("body: %s", issue.Body)
			fmt.Printf("created_at: %s\n", issue.CreatedAt)
			fmt.Printf("labels: %v\n", issue.Labels)
		}
	}
	if *readIssue { // read issue flag

		err := checkReadIssueFlags()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		issue, err := Read()
		if err != nil {
			fmt.Errorf("Read: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("repo: %s\n", issue.Repo)
		fmt.Printf("title: %s\n", issue.Title)
		fmt.Printf("body: %s", issue.Body)
		fmt.Printf("created_at: %s\n", issue.CreatedAt)
		fmt.Printf("labels: %v\n", issue.Labels)
	}

	if *closeIssue { // close issue flag
		err := Close() // close issue
		if err != nil {
			fmt.Errorf("Close: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("issue: %s is closed\n", issueLink())
	}
}
func checkReadIssueFlags() error {
	switch {
	case len(*collab) == 0:
		return fmt.Errorf("collab flag are not set")
	case len(*repo) == 0:
		return fmt.Errorf("repo flag are not set")
	case len(*owner) == 0:
		return fmt.Errorf("owner flag are not set")
	case len(*token) == 0:
		return fmt.Errorf("token flag are not set")
	case *issueNum == 0:
		return fmt.Errorf("issue number flag are not set")
	default:
		return nil
	}
}
func checkCreateUpdateFlags() error {
	switch {
	case len(*collab) == 0:
		return fmt.Errorf("collab flag are not set")
	case len(*repo) == 0:
		return fmt.Errorf("repo flag are not set")
	case len(*owner) == 0:
		return fmt.Errorf("owner flag are not set")
	case len(*token) == 0:
		return fmt.Errorf("token flag are not set")
	case len(*title) == 0:
		return fmt.Errorf("title flag are not set")
	default:
		return nil
	}
}

// SetHeaders
// Set the necessary http-headers.
// Устанавливаем необходимые http-заголовки для issue.Issue
func SetHeaders(req *http.Request) {
	//req.Header.Set("Accept", "application/vnd.github+json")

	req.Header.Set("Accept", "application/json") // create
	//req.Header.Set("Accept", "application/vnd.github+json")
	//application/vnd.github+json
	req.Header.Set("Authorization", "Bearer "+*token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")
}

// method PATCH
// /repos/{owner}/{repo}/issues/{issue_number}
func updateIssueLink() string {
	// /repos/{owner}/{repo}/issues/{issue_number}
	num := strconv.Itoa(*issueNum)
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", *owner, *repo, num)
}

// issuesLink
// Get link for all github user issues
// Получаем ссылку для всех issue репозитория
func issuesLink() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", *owner, *repo)
}

// issueLink
// Get link for one issue
// Получаем ссылку для одной issue
func issueLink() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", *owner, *repo, *issueNum)
}

// getLabels
// Splits a string into substrings and returns a slice of the substrings.
// Разделяет строку на подстроки, и возвращает срез подстрок.
func getLabels() []string {
	var trimmed string
	var labs []string

	if len(*labels) > 0 {
		trimmed = strings.TrimSpace(*labels) // clear spaces from the beginning and end of the line (очищаем от пробелов начало и конец строки)

		if strings.Contains(trimmed, ",") {
			labs = strings.Split(trimmed, ",")
		} else if strings.Contains(trimmed, " ") {
			labs = strings.Split(trimmed, " ")
		} else {
			labs = []string{0: trimmed}
		}
	}

	if len(labs) > 0 {
		return labs
	} else {
		return nil
	}
}
