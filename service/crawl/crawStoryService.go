package crawl

import (
	"fmt"
	"web/models"
	"web/store"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

type CrawlStoryService struct {
	Query     crawlStoryFrom
	ListStory []models.Story
}

type crawlStoryFrom struct {
	UrlVisit      string `form:"url_visit"`
	UrlDetail     string `form:"url_detail"`
	StartPaginate int    `form:"start_paginate"`
	StopPaginate  int    `form:"stop_paginate"`
	TitleStory    string `form:"story_title"`
	Auth          string `form:"auth"`
	CrawlFrame    string `form:"crawl_frame"`
}

func (cs *CrawlStoryService) Handle(ctx *gin.Context) {
	ctx.BindQuery(&cs.Query)
	db := store.ConnectDB()

	go cs.CrawlAll(db)

}

func (cs *CrawlStoryService) prepareCrawl(page int) {
	crawl := colly.NewCollector()

	crawl.OnHTML(cs.Query.CrawlFrame, func(e *colly.HTMLElement) {
		e.ForEach(".row", func(i int, el *colly.HTMLElement) {
			var story models.Story

			story.Name = el.ChildText(cs.Query.TitleStory)
			story.Auth = el.ChildText(cs.Query.Auth)
			story.AllowCrawl = true
			story.Status = 0

			crawlDescription := colly.NewCollector()
			crawlDescription.OnHTML(".desc-text", func(h *colly.HTMLElement) {
				story.Description = h.Text
			})

			crawlDescription.Visit(el.ChildAttr(".truyen-title a", "href"))

			if story.Name != "" && story.Auth != "" {
				cs.ListStory = append(cs.ListStory, story)
			}
		})
	})
	url := fmt.Sprintf(cs.Query.UrlVisit+"&page=%v", page)
	fmt.Println(url)
	crawl.Visit(url)
}

func (cs *CrawlStoryService) CrawlAll(db *gorm.DB) {
	for i := cs.Query.StartPaginate; i < cs.Query.StopPaginate; i++ {
		cs.prepareCrawl(i)
		db.CreateInBatches(cs.ListStory, 100)
		cs.ListStory = []models.Story{}
	}

}
