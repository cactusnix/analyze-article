package utils

import (
  "github.com/gocolly/colly"
  log "github.com/sirupsen/logrus"
  "strconv"
)

func GrabZhihu() {
  index := 266809594
  // 实例化一个收集器
  c := colly.NewCollector()
  // Add callbacks to a collector
  // Called before a request
  c.OnRequest(func(r *colly.Request) {
    log.Println("Visiting: ", r.URL)
  })
  // Called if error occured during the request
  c.OnError(func(r *colly.Response, err error) {
    if err != nil {
      log.Println("Something went wrong", err)
      index++
      r.Request.Visit("https://www.zhihu.com/api/v4/questions/" + strconv.Itoa(index))
    } else {
      log.Println("request url is", r.Request)
    }
  })
  // Called after response received
  c.OnResponse(func(r *colly.Response) {
    log.Println("status: ", r.StatusCode)
    log.Println("Visited", string(r.Body))
  })
  // Called right after OnResponse if the received content is HTML
  c.OnHTML(".ContentItem-title", func(e *colly.HTMLElement) {
    log.Println("class ContentItem-title", e)
  })
  url := "https://www.zhihu.com/api/v4/questions/" + strconv.Itoa(index)
  log.Println(url)
  c.Visit(url)
}
