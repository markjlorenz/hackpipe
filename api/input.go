package api

import (
  "io"
  "crypto/tls"
  "net/http"
  "fmt"
)

type Input struct {
  client *http.Client
  req *http.Request
}

type InputOpts struct {
  Path string
  Auth string
  Host string
  Yolo bool
}

func NewInput(opts *InputOpts) (*Input) {
  path := opts.Path
  auth := opts.Auth
  host := opts.Host
  yolo := opts.Yolo

  req, err := http.NewRequest("POST", "https://"+host+path, nil)
  if err != nil { fmt.Println(req); panic(err) }

  req.Header.Add("Authorization", auth)

  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: yolo},
  }

  return &Input{
    client: &http.Client{Transport: tr},
    req: req,
  }
}

func (i *Input)Write(body io.ReadCloser) io.ReadCloser{
  i.req.Body = body
  resp, err := i.client.Do(i.req)
  if err != nil { fmt.Println(resp); panic(err) }

  return resp.Body
}
