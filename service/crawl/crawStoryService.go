package crawl

import (
	"fmt"
	"strings"

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
		html, _ := e.DOM.Html()

		fmt.Printf("%+v", html)
		e.ForEach("div", func(_ int, el *colly.HTMLElement) {
			text := strings.TrimSpace(e.Text)
			fmt.Println(text)
		})
		// e.ForEach("div", func(i int, h *colly.HTMLElement) {
		// 	fmt.Printf("%+v: %+v", i, h.Text)
		// })
	})

	crawl.Visit(cs.Query.UrlVisit)
}
