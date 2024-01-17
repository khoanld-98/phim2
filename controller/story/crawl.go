package story

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Crawl(ctx *gin.Context) {
	crawler := colly.NewCollector()
	crawler.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting\n", r.URL)
	})

	// crawler.OnError(func(_ *colly.Response, err error) {
	// 	fmt.Printf("Something went wrong: ", err)
	// })

	ctx.JSON(200, gin.H{
		"message": 123,
	})

	crawler.OnHTML("#chapter-content", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Printf("===%+v---\n", e.Text)
	})
	crawler.Visit("https://dtruyen.com/thap-nien-60-xuyen-thanh-nong-nu-hai-thuoc-lam-giau/chuong-1_5124423.html")
	// crawler.OnScraped(func(r *colly.Response) {
	// 	fmt.Println(r.Request.URL, " scraped!")
	// })

}
