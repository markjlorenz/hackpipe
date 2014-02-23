package api

type Opts struct {
  Path    string
  Auth    string
  Headers map[string]string
  Host    string
  Query   string
  Method  string
  Yolo    bool
}
