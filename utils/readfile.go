package utils

import (
  "os"
  log "github.com/sirupsen/logrus"
  "bufio"
  "io"
  "strings"
  "io/ioutil"
  "encoding/json"
)

type Analysis struct {
  Language []string
  Framework []string
  ReadingWriting []string
}

// ReadFileByLine
func ReadFileByLine(path string) []string {
  var readResult []string
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  // 读取停用词表，封装成slice
  br := bufio.NewReader(file)
  for {
    // ReadString读取到分隔符，返回分隔符之前的数据包括分隔符
    exclude, err := br.ReadString('\n')
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    // 去除换行符
    readResult = append(readResult, strings.Trim(exclude, "\n"))
  }
  return readResult
}

// ReadFileAsString ...
func ReadFileAsString(path string) string {
  var readResult string
  content, err := ioutil.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  readResult = strings.Replace(string(content), "\n", "", -1)
  readResult = strings.Replace(readResult, "\t", "", -1)
  readResult = strings.Replace(readResult, "\r", "", -1)
  readResult = strings.Replace(readResult, "\v", "", -1)
  readResult = strings.Replace(readResult, "\f", "", -1)
  readResult = strings.Replace(readResult, " ", "", -1)
  readResult = strings.ToLower(readResult)
  return readResult
}

// ReadJson ...
func ReadJson(path string) Analysis {
  var readResult Analysis
  content, err := ioutil.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  err1 := json.Unmarshal(content, &readResult)
  if err1 != nil {
    log.Fatal(err)
  }
  return readResult
}
