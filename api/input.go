package api

import (
  "io"
  "crypto/tls"
  "net/http"
  "fmt"
)

type Input struct {
  client *http.Client
  Request *http.Request
}

func NewInput(opts *Opts) (*Input) {
  path  := opts.Path
  auth  := opts.Auth
  host  := opts.Host
  yolo  := opts.Yolo
  head  := opts.Headers
  meth  := opts.Method
  query := opts.Query

  req, err := http.NewRequest(meth, "https://"+host+path, nil)
  if err != nil { fmt.Println(req); panic(err) }

  req.Header.Add("Authorization", auth)

  for key, value := range head {
    req.Header.Add(key, value)
  }

  req.URL.RawQuery = query

  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: yolo},
  }

  return &Input{
    client: &http.Client{Transport: tr},
    Request: req,
  }
}

func (i *Input)Write(body io.ReadCloser) io.ReadCloser{
  i.Request.Body = body
  resp, err := i.client.Do(i.Request)
  if err != nil { fmt.Println(resp); panic(err) }

  return resp.Body
}
