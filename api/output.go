package api

import (
  "crypto/tls"
  "net/http"
  "bufio"
  "fmt"
)

type Output struct {
  *bufio.Reader
}

func NewOutput(opts *Opts) *Output {
  path := opts.Path
  auth := opts.Auth
  host := opts.Host
  yolo := opts.Yolo
  head := opts.Headers

  req, err := http.NewRequest("GET", "https://"+host+path, nil)
  if err != nil { fmt.Println(req); panic(err) }

  req.Header.Add("Authorization", auth)

  for key, value := range head {
    req.Header.Add(key, value)
  }

  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: yolo},
  }

  client := &http.Client{Transport: tr}

  resp, err := client.Do(req)
  reader := bufio.NewReader(resp.Body)

  return &Output{
    Reader: reader,
  }
}
