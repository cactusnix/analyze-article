package utils

import (
  "github.com/gocolly/colly"
  log "github.com/sirupsen/logrus"
)

func GrabZhihu() {
  // 实例化一个收集器
  c := colly.NewCollector()
  // Add callbacks to a collector
  // Called before a request
  c.OnRequest(func(r *colly.Request) {
    log.Println("Visiting: ", r.URL)
  })
  // Called if error occured during the request
  c.OnError(func(_ *colly.Response, err error) {
    log.Println("Something went wrong:", err)
  })
  // Called after response received
  c.OnResponse(func(r *colly.Response) {
    log.Println("Visited", r.Request.URL)
  })
  // Called right after OnResponse if the received content is HTML
  c.OnHTML(".ContentItem-title", func(e *colly.HTMLElement) {
    log.Println("class ContentItem-title", e)
  })
  c.Visit("https://www.zhihu.com/")
}
