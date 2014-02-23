package filter

import (
  "os/exec"
  "io/ioutil"
  "bytes"
  "fmt"
  "strings"
  "net/http"
  "os"
)

const NL string = "\n"

type Filtered struct {
  bytes.Buffer
}
func (f *Filtered) Close() error { return nil }

type Filter struct {
  runner    string
  script    string
  queryFile string
}

type NullFilter struct { }
func (f *NullFilter) Filter(raw *bytes.Buffer, filtered *Filtered, req *http.Request) {
  filtered.ReadFrom(raw)
}

type Filterable interface {
  Filter(raw *bytes.Buffer, filtered *Filtered, req *http.Request)
}

func NewFilter(runner, script string) Filterable {
  // if we can't filter, don't try.
  if runner == "" || script == "" {
    return &NullFilter{}
  }

  return &Filter {
    runner: runner,
    script: script,
  }
}

func (f *Filter) Filter(raw *bytes.Buffer, filtered *Filtered, req *http.Request) {
  filename := f.writeScriptFile()

  commands := strings.Fields(f.runner)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd := exec.Command(commands[0], args...)
  cmd.Stdin = raw

  if req != nil {
    f.setupSpecialFiles(cmd, req)
  }

  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(string(scripted)); panic(err) }

  if req !=nil {
    f.updateQuery(req)
  }

  noTrailing := strings.TrimSuffix(string(scripted), NL)

  fmt.Fprint(filtered, noTrailing)
  return
}

func (f *Filter) writeScriptFile() string {
  tmp, err := ioutil.TempFile("", "hackpipe:")
  if err != nil { panic(err) }

  defer tmp.Close()
  _, err  = tmp.WriteString(f.script)
  if err != nil { panic(err) }

  return tmp.Name()
}

func (f *Filter) setupSpecialFiles(cmd *exec.Cmd, req *http.Request) {
  query, err := ioutil.TempFile("", "hackpipe:")
  if err != nil { panic(err) }
  defer query.Close()

  f.queryFile = query.Name()

  _, err = query.WriteString(req.URL.RawQuery)
  if err != nil { panic(err) }

  // make the query file available
  // The just come in as FD3 to FD<3+num files>
  // cmd.ExtraFiles = []*os.File{
  //   query,
  // }

  // set convenience ENV to the descriptors
  cmd.Env = []string{
    "QUERY="+query.Name(),
  }
}

func (f *Filter) updateQuery(req *http.Request) {
  x := f.queryFile
  queryFile, err :=  os.Open(x)
  if err != nil { panic(err) }

  queryBytes, err :=  ioutil.ReadAll(queryFile)
  if err != nil { panic(err) }

  req.URL.RawQuery = string(queryBytes)
}
