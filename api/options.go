package api

type Opts struct {
  Scheme  string
  Path    string
  Auth    string
  Headers map[string]string
  Host    string
  Query   string
  Method  string
  Yolo    bool
}
