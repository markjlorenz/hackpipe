package api

import (
  "io"
  "crypto/tls"
  "net/http"
  "fmt"
)

func NewInput(body io.Reader) io.ReadCloser {
  path := "/gists"
  auth := "Basic Nzc5ZDQ3Mzg2MjI0ZmZkODBlYTI5OWRkMzNjYzllYmZmYWYxNGMyZDp4LW9hdXRoLWJhc2lj"
  host := "api.github.com"
  yolo := false

  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: yolo},
  }
  client := &http.Client{Transport: tr}

  req, err := http.NewRequest("POST", "https://"+host+path, body)
  if err != nil { fmt.Println(req); panic(err) }

  req.Header.Add("Authorization", auth)

  resp, err := client.Do(req)
  if err != nil { fmt.Println(resp); panic(err) }

  return resp.Body
}
