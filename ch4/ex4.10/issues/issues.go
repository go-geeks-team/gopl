// Exercise 4.10
//
// Пример сканирования issues поданные менее месяца назад (только первые 10 страниц):
//
// $ ./issues -start_page=1 -end_page=10 -period=less_month_ago -query=bug
//
// Подсказка по флагам:
//    -start_page=1          — начинаем сканирование с первой страницы
//    -end_page=10           — заканчиваем сканирование на 10 странице (включительно)
//    -period=less_month_ago — выводим результаты поданные менее месяца назад
//    -query="bug"           — поиск issues содержащим в title строку "bug"

package main

import (
	"flag"
	"fmt"
	"issues/github"
	"os"
	"strings"
	"time"
)

const (
	LessMonthAgo = "less_month_ago" // менее месяца назад
	LessYearAgo  = "less_year_ago"  // менее года назад
	MoreYearAgo  = "more_year_ago"  // более года назад
)
const (
	DayHours   = 24
	MonthHours = 730 // MonthHours Количество часов в одном месяце (приблизительно)
)

var (
	startPage = flag.Int("start_page", 1, "Default start page")                                                         // старт сканирования (значение по-умолчанию 1)
	endPage   = flag.Int("end_page", 10, "Default end page")                                                            // заканчиваем сканирование (значение по-умолчанию 10)
	perPage   = flag.Int("per_page", 10, "Rows count per page")                                                         // количество строк на страницу
	query     = flag.String("query", "", "Query string")                                                                // строка поиска
	period    = flag.String("period", "", "Possible flag values: -period={less_month_ago|less_year_ago|more_year_ago}") // период: менее_месяца назад, менее_года_назад, более_года_назад
)

var (
	issues []*github.Issue // собираем в issues отфильтрованные по периуду (см. флаг period)
	totCnt int             // total issues count
)

var threadSleep = time.Duration(7000) // prevent request blocking
var errsCnt int

func main() {
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) { // проходимся по флагам period и query проверяя их корректность
		if f.Name == "period" && f.Value.String() != MoreYearAgo && f.Value.String() != LessYearAgo && f.Value.String() != LessMonthAgo {
			head := "Flag period is empty\nPossible values: period={less_month_ago | less_year_ago | more_year_ago}\n"
			tail := "Example1: -period=less_month_ago\nExample2: -period=less_year_ago\nExample3: -period=more_year_ago\n"
			fmt.Print(head + tail)
			errsCnt++
		}
		if f.Name == "query" && len(f.Value.String()) == 0 {
			head := "\nQuery flag is empty\n"
			tail := "Example:\n ./issues -period=less_month_ago -query=\"memory leak\"\n"
			fmt.Print(head + tail)
			errsCnt++
		}

	})
	if errsCnt > 0 {
		os.Exit(1)
	}

	fmt.Print("start scanning...\n")

	printAsTable(collectIssues())
	// Scan issues submitted a month ago (only the first five pages):
	//
	// Flags hint:
	//    -start_page=1           — start scanning from the first page
	//    -end_page=10            — finish scanning on page 10 (inclusive)
	//    -period=less_month_ago  — display results submitted "less than a month ago"
	//    -query="bug"            — search for issues containing the string "bug" in the title
	//
	// $ ./issues -start_page=1 -end_page=5 -period=less_month_ago -query=bug
}

// printHorLine
// Print a horizontal line.
// Распечатываем горизонтальную линию.
func printHorLine(width int) {
	var sb strings.Builder
	for i := 0; i < width; i++ {
		sb.WriteString("-")
	}
	fmt.Printf("%s\n", sb.String())
}

// printAsTable
// Print pretty pseudo table.
// Распечатываем аккуратную псевдо-таблицу.
func printAsTable(issues []*github.Issue) {
	fmt.Printf("total issues count: %d\n", totCnt)
	fmt.Print("\n")
	printHorLine(100)
	fmt.Printf("| %-4s | %9.9s | %9.9s | %9.9s | %.55s \n", "Count", "Days ago", "Date", "Login", "Title") // print header
	printHorLine(100)
	for i, issue := range issues {
		fmt.Printf("| %-4d | %9.9d | %9.9s | %9.9s | %.55s \n", i+1, issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
	}
	printHorLine(100)

	fmt.Printf("start page: %d\n", *startPage)
	fmt.Printf("end page: %d\n", *endPage)
	fmt.Printf("per page: %d\n", *perPage)
	fmt.Printf("total issues count: %d\n", totCnt)
}

// Search Issues:
//
// submitted less than a month ago
// submitted less than one year ago
// submitted more than a year ago
//
// Поиск Issues:
//
//	поданные менее месяца назад
//	поданные менее года назад
//	поданные более года назад
func searchIssues(issues []*github.Issue, page int) ([]*github.Issue, error) {

	// wait 7 sec for preventing request blocking
	time.Sleep(time.Millisecond * threadSleep) // усыпляем текущий поток на 7 секунд (предотвращаем блокировку).

	fmt.Printf("scanning page %-4d", page)
	isr, err := github.SearchIssues2(strings.Split(*query, " "), page, *perPage)
	if err != nil {
		return nil, err
	}
	cnt := 0
	for _, issue := range isr.Items {

		start := issue.CreatedAt

		// submitted less than a month ago
		// отправлено менее месяца назад
		if *period == LessMonthAgo && submittedLessThanMonthAgo(start) {
			issues = append(issues, issue)
			cnt++ // increase counter
		}

		// submitted less than one year ago
		// поданные менее года назад
		if *period == LessYearAgo && submittedLessThanOneYearAgo(start) {
			issues = append(issues, issue)
			cnt++ // increase counter
		}

		// submitted more than a year ago
		// поданные более года назад
		if *period == MoreYearAgo && submittedMoreThanYearAgo(start) {
			issues = append(issues, issue)
			cnt++ // increase counter
		}

	}
	totCnt += cnt
	fmt.Printf("\tfound %d issues\n", cnt)

	return issues, nil
}

// Returns a slice of issues with elements added less than a month ago, less than a year ago, and more than a year ago
// Возвращает срез issues элементами добаленные меньше месяца назад, меньше года назад и более года назад
func collectIssues() []*github.Issue {

	fmt.Print("# Search Issues Submitted")

	switch *period {
	case LessMonthAgo:
		fmt.Printf(" a less month ago\n")
	case LessYearAgo:
		fmt.Printf(" a less year ago\n")
	case MoreYearAgo:
		fmt.Printf(" a more year ago\n")
	}

	fmt.Printf("# Search word(s): %s\n", *query)

	for curPage := *startPage; curPage <= *endPage; curPage++ {
		issuesPage, err := searchIssues(nil, curPage)
		if err != nil {
			break
		}
		if len(issuesPage) > 0 {
			issues = append(issues, issuesPage...)
		}
	}
	return issues
}

// submittedLessThanMonthAgo issues отправленные менее месяца назад
func submittedLessThanMonthAgo(start time.Time) bool {
	end := time.Now()
	return int(end.Sub(start).Hours()) < DayHours*31 && int(end.Sub(start).Hours()) > DayHours*7
}

// submittedLessThanOneYearAgo issues отправленные менее года назад
func submittedLessThanOneYearAgo(start time.Time) bool {
	end := time.Now()
	return int(end.Sub(start).Hours()) >= MonthHours*7 && int(end.Sub(start).Hours()) < MonthHours*11
}

// submittedMoreThanYearAgo issues отправленные более года назад
func submittedMoreThanYearAgo(start time.Time) bool {
	end := time.Now()
	return int(end.Sub(start).Hours()) > MonthHours*12 && int(end.Sub(start).Hours()) < MonthHours*24
}
