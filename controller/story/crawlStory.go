package story

import (
	"web/service/crawl"

	"github.com/gin-gonic/gin"
)

func ScrawlStory(ctx *gin.Context) {
	var crawlStoryService crawl.CrawlStoryService

	crawlStoryService.Handle(ctx)
}
