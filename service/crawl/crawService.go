package crawl

import (
	"fmt"
	"web/models"
	"web/store"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type CrawlService struct {
	query        crawlForm
	listContent  []models.Chap
	linkPaginate []string
	NextPage     string
	Story        models.Story
}

type crawlForm struct {
	Name       string `form:"name"`
	Url        string `form:"url"`
	NameSeries string `form:"name_series"`
	NextPage   string `form:"next_page"`
	Content    string `form:"content"`
}

func (cs *CrawlService) Handle(ctx *gin.Context) {
	db := store.ConnectDB()
	ctx.BindQuery(&cs.query)

	story := models.Story{
		Name: cs.query.Name,
	}

	db.Create(&story)
	cs.Story = story

	for {
		check := cs.prepareCrawl()
		if check == false {
			break
		}
	}

	db.CreateInBatches(cs.listContent, 100)
}

func (cs *CrawlService) prepareCrawl() bool {
	crawler := colly.NewCollector()
	var content models.Chap
	// name series
	crawler.OnHTML(cs.query.NameSeries, func(e *colly.HTMLElement) {
		content.Name = e.Text
	})

	// url page
	crawler.OnHTML(cs.query.NextPage, func(e *colly.HTMLElement) {
		cs.NextPage = e.Attr("href")
	})

	// content of series
	crawler.OnHTML(cs.query.Content, func(e *colly.HTMLElement) {
		content.Content = e.Text
	})

	crawler.Visit(cs.query.Url)
	content.MovieId = int(cs.Story.ID)

	fmt.Printf("get Content: %+v\n", cs.query.Url)
	cs.query.Url = cs.NextPage

	cs.listContent = append(cs.listContent, content)

	if cs.NextPage == "javascript:void(0)" || cs.NextPage == "#" {
		return false
	}

	return true
}
