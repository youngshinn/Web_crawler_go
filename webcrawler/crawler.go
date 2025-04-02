package webcrawler

import (
	"fmt"
	"goproject/database"
	"log"
	"sync"

	"github.com/gocolly/colly/v2"
)

const baseURL = "https://search.naver.com/search.naver?where=news&query=%s&start=%d"

func CrawlPage(keyword string, page int, wg *sync.WaitGroup, results chan<- database.News) {
	defer wg.Done()

	c := colly.NewCollector()

	c.OnHTML(".news_area", func(e *colly.HTMLElement) {
		title := e.ChildText(".news_tit")
		link := e.ChildAttr(".news_tit", "href")
		results <- database.News{
			Title:   title,
			Link:    link,
			Keyword: keyword,
		}
	})

	url := fmt.Sprintf(baseURL, keyword, (page-1)*10+1)
	err := c.Visit(url)
	if err != nil {
		log.Println("크롤링 실패:", err)
	}
}
