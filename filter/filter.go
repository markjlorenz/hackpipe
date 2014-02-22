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

type Filtered struct {
  bytes.Buffer
}
func (f *Filtered) Close() error { return nil }

type Filter struct {
  command string
  script  string
}

type NullFilter struct { }
func (f *NullFilter) Filter(raw *bytes.Buffer, filtered *Filtered) {
  filtered.ReadFrom(raw)
}

type Filterable interface {
  Filter(raw *bytes.Buffer, filtered *Filtered)
}

func NewFilter(command, script string) Filterable {
  // if we can't filter, don't try.
  if command == "" || script == "" {
    return &NullFilter{}
  }

  return &Filter {
    command: command,
    script:  script,
  }
}

func (f *Filter) Filter(raw *bytes.Buffer, filtered *Filtered) {
  filename := f.writeScriptFile()

  commands := strings.Fields(f.command)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd := exec.Command(commands[0], args...)
  cmd.Stdin = raw

  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(string(scripted)); panic(err) }

  fmt.Fprint(filtered, string(scripted))
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
