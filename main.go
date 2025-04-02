// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"goproject/database"
	"goproject/webcrawler"
)

func main() {
	database.ConnectDB()

	if len(os.Args) < 3 {
		fmt.Println("사용법: go run main.go <키워드> <페이지 수>")
		return
	}

	keyword := os.Args[1]
	numPages, err := strconv.Atoi(os.Args[2])
	if err != nil || numPages < 1 {
		fmt.Println("올바른 페이지 수를 입력하세요.")
		return
	}

	var wg sync.WaitGroup
	results := make(chan database.News, numPages*10)

	for page := 1; page <= numPages; page++ {
		wg.Add(1)
		go webcrawler.CrawlPage(keyword, page, &wg, results)
	}

	// 저장 쓰레드 실행
	go func() {
		wg.Wait()
		close(results)
	}()

	for news := range results {
		database.SaveNews(news.Title, news.Link, news.Keyword)
	}
}
