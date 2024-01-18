package story

import (
	"web/service/crawl"

	"github.com/gin-gonic/gin"
)

func Crawl(ctx *gin.Context) {
	var crawService crawl.CrawlService

	crawService.Handle(ctx)
}
