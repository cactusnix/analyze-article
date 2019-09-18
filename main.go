package main

import (
  "github.com/yanyiwu/gojieba"
  log "github.com/sirupsen/logrus"
  "alyatc/utils"
)

func main() {
  // nil 切片的长度和容量为 0 且没有底层数组。
  var words []string
  var excludes []string
  // make只用于slice、map、
  wordCount := make(map[string]int)
  x := gojieba.NewJieba()
  defer x.Free()
  use_hmm := true
  content := utils.ReadFileAsString("data/analyse_articles/简书一年，我出版了一本书.txt")
  words = x.Cut(content, use_hmm)
  // log.Info(words)
  excludes = utils.ReadFileByLine("data/stopwords/中文停用词表.txt")
  // log.Info(excludes)
  // 去除停用词
  for i := 0; i < len(words); i++ {
    for j := 0; j < len(excludes); j++ {
      if words[i] == excludes[j] {
        words = append(words[:i], words[i+1:]...)
        i--
        // use 'break' to fix the error index out of the range
        break
      }
    }
  }
  // 统计词频
  for _, word := range words {
    if wordCount[word] != 0 {
      wordCount[word]++
    } else {
      wordCount[word] = 1
    }
  }
  // log.Info(wordCount)
  // log.Info(len(wordCount))
  wordResult := utils.RankByValue(wordCount)[0:30]
  log.Info(wordResult)
  // load the ref words
  analysis := utils.ReadJson("data/analyse_reference/analysis.json")
  lanWords := analysis.Language
  lanCount := utils.Count(wordResult, lanWords)
  frameWords := analysis.Framework
  frameCount := utils.Count(wordResult, frameWords)
  rwWords := analysis.ReadingWriting
  log.Info(rwWords)
  rwCount := utils.Count(wordResult, rwWords)
  log.Info("文章分析结果：")
  log.Println("编程语言讨论占比", lanCount, len(wordCount), float64(lanCount)/float64(len(wordCount)))
  log.Println("编程框架讨论占比", frameCount, len(wordCount), float64(frameCount)/float64(len(wordCount)))
  log.Println("读书写作讨论占比", rwCount, len(wordCount), float64(rwCount)/float64(len(wordCount)))
}
