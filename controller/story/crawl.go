package story

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Crawl(ctx *gin.Context) {
	crawler := colly.NewCollector()
	ctx.JSON(200, gin.H{
		"message": 123,
	})
	err := crawler.Visit("https://dtruyen.com/thap-nien-60-xuyen-thanh-nong-nu-hai-thuoc-lam-giau/chuong-2_5124424.html")
	fmt.Printf("======================Visiting: %+v======================", err)
	fmt.Printf("======================Visiting: %+v======================", crawler)
	crawler.OnRequest(func(r *colly.Request) {
		fmt.Printf("======================Visiting: %+v======================", r.URL)
	})

	// crawler.OnError(func(_ *colly.Response, err error) {
	// 	fmt.Printf("Something went wrong: ", err)
	// })

	// crawler.OnResponse(func(r *colly.Response) {
	// 	fmt.Printf("Page visited: ", r.Request.URL)
	// })

	// crawler.OnHTML("a", func(e *colly.HTMLElement) {
	// 	// printing all URLs associated with the a links in the page
	// 	fmt.Printf("%+v", e)
	// })

	// crawler.OnScraped(func(r *colly.Response) {
	// 	fmt.Println(r.Request.URL, " scraped!")
	// })

}
