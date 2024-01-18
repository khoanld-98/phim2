package crawl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type CrawlService struct {
	query        crawlForm
	listContent  []content
	linkPaginate []string
}

type crawlForm struct {
	Name       string `form:"Name"`
	Url        string `form:"url"`
	NameSeries string `form:"name_series"`
	NextPage   string `form:"next_page"`
	Content    string `form:"content"`
}

type content struct {
	Name    string
	Content string
}

func (cs *CrawlService) Handle(ctx *gin.Context) {
	ctx.BindQuery(&cs.query)

	for {
		check := cs.prepareCrawl()
		if check == false {
			break
		}
	}

	fmt.Printf("%+v", cs.listContent)
}

func (cs *CrawlService) prepareCrawl() bool {
	crawler := colly.NewCollector()
	var content content
	// name series
	crawler.OnHTML(cs.query.NameSeries, func(e *colly.HTMLElement) {
		content.Name = e.Text
	})

	// url page
	crawler.OnHTML(cs.query.NextPage, func(e *colly.HTMLElement) {
		fmt.Printf("get Content in page1: %+v\n", e.Attr("href"))
		if e.Attr("href") != "javascript:void(0)" || e.Attr("href") != "#" {
			fmt.Printf("get Content in page2: %+v\n", e.Attr("href"))
			cs.query.Url = e.Attr("href")
		}

		cs.query.NextPage = e.Attr("href")
	})

	// content of series
	crawler.OnHTML(cs.query.Content, func(e *colly.HTMLElement) {
		content.Content = e.Text
	})

	if cs.query.Url == "" {
		return false
	}

	crawler.Visit(cs.query.Url)
	// fmt.Printf("get Content: %+v", content)
	fmt.Printf("get Content: %+v\n", cs.query.Url)
	cs.listContent = append(cs.listContent, content)

	return true
}
