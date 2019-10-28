package utils

import (
  "github.com/gocolly/colly"
  log "github.com/sirupsen/logrus"
  "os"
  "bufio"
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
  c.OnError(func(r *colly.Response, err error) {
    if err != nil {
      log.Println("Something went wrong", err)
      r.Request.Visit("https://www.zhihu.com/api/v4/questions/")
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
  url := "https://www.zhihu.com/api/v4/questions/"
  log.Println(url)
  c.Visit(url)
}

func GrabJianShu() {
  name := ""
  c := colly.NewCollector(
    colly.Async(true),
  )
  c.OnRequest(func(r *colly.Request) {
    log.Println("Visiting: ", r.URL)
  })
  c.OnError(func(r *colly.Response, err error) {
    if err != nil {
      log.Fatal("Error: ", err)
    }
  })
  c.OnResponse(func(r *colly.Response) {
    log.Println("status: ", r.StatusCode)
  })
  articleController := colly.NewCollector(
    colly.Async(true),
  )
  articleController.OnRequest(func(r *colly.Request) {
    log.Println("Visiting: ", r.URL)
  })
  articleController.OnError(func(r *colly.Response, err error) {
    if err != nil {
      log.Fatal("Error: ", err)
    }
  })
  articleController.OnResponse(func(r *colly.Response) {
    log.Println("status: ", r.StatusCode)
  })
  c.OnHTML("a[href].title", func(e *colly.HTMLElement) {
    log.Println("article title", e.Text, e.Attr("href"))
    if e.Attr("href") != "" {
      name = e.Text + ".txt"
      articleController.Visit("https://www.jianshu.com" + e.Attr("href"))
      articleController.Wait()
    }
  })
  articleController.OnHTML("p", func(e *colly.HTMLElement) {
    log.Println(e.Text)
    openThenWrite(name, e.Text)
  })
  url := "https://www.jianshu.com"
  c.Visit(url)
  c.Wait()
}

func openThenWrite(fileName string, content string) {
  filePath := "data/analyze_articles/" + fileName
  file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  writer := bufio.NewWriter(file)
  writer.WriteString(content + "\n")
  writer.Flush()
}