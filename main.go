package main

import (
  log "github.com/sirupsen/logrus"
  "io/ioutil"
  "alyatc/app"
  // "alyatc/utils"
)

func main() {
  // utils.GrabJianShu()
  files, err := ioutil.ReadDir("data/analyze_articles/")
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    app.Analyze(file.Name())
  }
}
