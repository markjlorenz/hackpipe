package filter

import (
  "os/exec"
  "io/ioutil"
  "bytes"
  "fmt"
  "strings"
)

const NL string = "\n"

type Filtered struct {
  bytes.Buffer
}
func (f *Filtered) Close() error { return nil }

type Filter struct {
  runner string
  script string
}

type NullFilter struct { }
func (f *NullFilter) Filter(raw *bytes.Buffer, filtered *Filtered) {
  filtered.ReadFrom(raw)
}

type Filterable interface {
  Filter(raw *bytes.Buffer, filtered *Filtered)
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

func (f *Filter) Filter(raw *bytes.Buffer, filtered *Filtered) {
  filename := f.writeScriptFile()

  commands := strings.Fields(f.runner)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd := exec.Command(commands[0], args...)
  cmd.Stdin = raw

  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(string(scripted)); panic(err) }

  noTrailing := strings.TrimSuffix(string(scripted), NL)

  fmt.Fprint(filtered, noTrailing)
  return
}

func (f *Filter) writeScriptFile() string {
  tmp, err := ioutil.TempFile("/tmp", "hackpipe:")
  if err != nil { panic(err) }

  defer tmp.Close()
  _, err  = tmp.WriteString(f.script)
  if err != nil { panic(err) }

  return tmp.Name()
}
//
// func (f *Filter) setupSpecialFiles() {
//   tmp, err := ioutil.TempFile("/tmp", "hackpipe:")
// }
