package filter

import (
  "os"
  "os/exec"
  "bytes"
  "fmt"
  "time"
  "strings"
  "strconv"
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

func (f *Filter) writeScriptFile() (filename string) {
  filename = "/tmp/hackpipe:"+strconv.FormatInt(time.Now().UnixNano(), 10)
  file, _ := os.Create(filename)
  defer file.Close()
  _, err := file.WriteString(f.script)
  if err != nil { panic(err) }

  return
}
