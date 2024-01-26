package crawl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type CrawlStoryService struct {
	Query crawlStoryFrom
}

type crawlStoryFrom struct {
	UrlVisit      string `form:"url_visit"`
	UrlDetail     string `form:"url_detail"`
	StartPaginate int    `form:"start_paginate"`
	StopPaginate  int    `form:"stop_paginate"`
	TitleStory    string `form:"story_title"`
	Auth          string `form:"auth"`
}

func (cs *CrawlStoryService) Handle(ctx *gin.Context) {
	ctx.BindQuery(&cs.Query)

	cs.prepareCrawl(1)
}

func (cs *CrawlStoryService) prepareCrawl(page int) {
	crawl := colly.NewCollector()

	crawl.OnHTML(cs.Query.Auth, func(e *colly.HTMLElement) {
		fmt.Printf("---------------%+v------------------", e.Text)
	})
}
