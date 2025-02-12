### Exercise 4.11:

Build a tool that lets users create, read, update, and delete GiHub issues from the command line,
invoking their preferred text editor when substantial text input is required.

---

### Упражнение 4.11:

Создайте инструмент, который позволит пользователю создавать, читать, обновлять и закрывать темы GitHub из 
командной строки, вызывая предпочитаемый пользователем текстовый редактор, когда требуется ввести текст значительного
объёма.

---

### Create Issue


Example Schema:
```bash
$ go test -v -create
  -collab="<GITHUB_COLLABORATOR>"
  -token="<COLLABORATOR_GITHUB_TOKEN>"
  -repo="<GITHUB_REPO>"
  -owner="<REPO_OWNER>"
  -title="<ISSUE_TITLE>"
  -body="<ISSUE_BODY>"
```

Example:
Run test:
```bash
$ go test -v -create  -collab="<GITHUB_COLLABORATOR>" -token="<COLLABORATOR_GITHUB_TOKEN>" -repo="testRepo1" -owner=unixlinuxgeek -title="NewIssue" -body="new issuev body"
=== RUN   TestCreate
    issue_test.go:111: state: 
    issue_test.go:112: repo: testRepo1
    issue_test.go:113: owner: unixlinuxgeek
    issue_test.go:114: title: NewIssue
    issue_test.go:115: body: new issue
    issue_test.go:116: labels: []
    issue_test.go:117: created_at: 0001-01-01 00:00:00 +0000 UTC
    issue_test.go:118: HTMLURL: 
    issue_test.go:119: number: 0
--- PASS: TestCreate (1.39s)
=== RUN   TestUpdate
--- PASS: TestUpdate (0.00s)
=== RUN   TestReadIssue
--- PASS: TestReadIssue (0.00s)
=== RUN   TestCloseIssue
--- PASS: TestCloseIssue (0.00s)
PASS
ok      issue   1.395s

```

---

### Update Issue

Example Schema:
```bash
$ go test -update
 -collab="<GITHUB_COLLABORATOR>"
 -token="<GITHUB_COLLABORATOR_TOKEN>"
 -repo="<GITHUB_REPO>"
 -owner="<REPO_OWNER>"
 -title="<ISSUE_TITLE>"
 -body="<ISSUE_BODY>"
 -issue_number=<GITHUB_ISSUE_NUMBER>
```

Example output:
```bash
$ go test -v -update -collab="<GITHUB_COLLABORATOR>" -token="<GITHUB_COLLABORATOR_TOKEN>" -repo="testRepo1" -owner=unixlinuxgeek issue_num=43
=== RUN   TestCreate
--- PASS: TestCreate (0.00s)
=== RUN   TestUpdate
    issue_test.go:111: state: 
    issue_test.go:112: repo: testRepo1
    issue_test.go:113: owner: unixlinuxgeek
    issue_test.go:114: title: 
    issue_test.go:115: body: 
    issue_test.go:116: labels: []
    issue_test.go:117: created_at: 0001-01-01 00:00:00 +0000 UTC
    issue_test.go:118: HTMLURL: 
    issue_test.go:119: number: 0
--- PASS: TestUpdate (1.36s)
=== RUN   TestReadIssue
--- PASS: TestReadIssue (0.00s)
=== RUN   TestCloseIssue
--- PASS: TestCloseIssue (0.00s)
PASS
ok      issue   1.364s
```
---

#### Read Issue

```bash
$ go test -v -read
  -collab="<GITHUB_COLLABORATOR>"
  -token="<COLLABORATOR_GITHUB_TOKEN>"
  -repo="<GITHUB_REPO>"
  -owner="<REPO_OWNER>"
  -issue_number=43
```

Unit test output:
```bash
$ go test -v -read -collab="<GITHUB_COLLABORATOR>" -token="<GITHUB_COLLABORATOR_TOKEN>" -repo="testRepo1" -owner="unixlinuxgeek" -issue_num=43 
=== RUN   TestCreate
--- PASS: TestCreate (0.00s)
=== RUN   TestUpdate
--- PASS: TestUpdate (0.00s)
=== RUN   TestReadIssue
    issue_test.go:111: state: open
    issue_test.go:112: repo: 
    issue_test.go:113: owner: 
    issue_test.go:114: title: New000
    issue_test.go:115: body: New000_body
    issue_test.go:116: labels: []
    issue_test.go:117: created_at: 2025-01-16 12:48:04 +0000 UTC
    issue_test.go:118: HTMLURL: https://github.com/unixlinuxgeek/testRepo1/issues/43
    issue_test.go:119: number: 43
--- PASS: TestReadIssue (1.31s)
=== RUN   TestCloseIssue
--- PASS: TestCloseIssue (0.00s)
PASS
ok      issue   1.315s
```

---

#### Close Issue

```bash
$ go test -v -close  
    -collab="<GITHUB_COLLABORATOR>" 
    -token="<GITHUB_COLLABORATOR_TOKEN>" 
    -repo="testRepo1" 
    -owner=unixlinuxgeek 
    -issue_number=43
```

Unit test output:
```bash
$ go test -v -close  -collab="<GITHUB_COLLABORATOR>" -token="<GITHUB_COLLABORATOR_TOKEN>" -repo="testRepo1" -owner=unixlinuxgeek -issue_num=43
=== RUN   TestCreate
--- PASS: TestCreate (0.00s)
=== RUN   TestUpdate
--- PASS: TestUpdate (0.00s)
=== RUN   TestReadIssue
--- PASS: TestReadIssue (0.00s)
=== RUN   TestCloseIssue
issue_test.go:111: state: closed
issue_test.go:112: repo: testRepo1
issue_test.go:113: owner: unixlinuxgeek
issue_test.go:114: title:
issue_test.go:115: body:
issue_test.go:116: labels: []
issue_test.go:117: created_at: 0001-01-01 00:00:00 +0000 UTC
issue_test.go:118: HTMLURL:
issue_test.go:119: number: 43
--- PASS: TestCloseIssue (1.26s)
PASS
ok      issue   1.263s
```
