package filter

import (
  "os/exec"
  "io/ioutil"
  "bytes"
  "net/http"
  "os"
)

type InputFilter struct {
  Filter
  queryFile string
  request   *http.Request
}

func NewInputFilter(runner, script string, req *http.Request) Filterable {
  // if we can't filter, don't try.
  if runner == "" || script == "" {
    return &NullFilter{}
  }

  return &InputFilter {
    Filter{
      runner:  runner,
      script:  script,
    },
    "",
    req,
  }
}

func (f *InputFilter) Run(raw *bytes.Buffer, filtered *Filtered) {
  cmd := f.getCommand(raw)

  f.setupSpecialFiles(cmd)

  scripted := f.runCommand(cmd)

  f.updateQuery()

  f.writeResult(scripted, filtered)
}

func (f *InputFilter) setupSpecialFiles(cmd *exec.Cmd) {
  query, err := ioutil.TempFile("", PREFIX)
  if err != nil { panic(err) }
  defer query.Close()

  f.queryFile = query.Name()

  _, err = query.WriteString(f.request.URL.RawQuery)
  if err != nil { panic(err) }

  cmd.Env = []string{
    "QUERY="+query.Name(),
  }
}

func (f *InputFilter) updateQuery() {
  x := f.queryFile
  queryFile, err :=  os.Open(x)
  if err != nil { panic(err) }

  queryBytes, err :=  ioutil.ReadAll(queryFile)
  if err != nil { panic(err) }

  f.request.URL.RawQuery = string(queryBytes)
}
