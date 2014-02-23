package filter

import (
  "os/exec"
  "io/ioutil"
  "bytes"
  "fmt"
  "strings"
)

const NL string = "\n"
const PREFIX string = "hackpipe:"

type Filtered struct {
  bytes.Buffer
}
func (f *Filtered) Close() error { return nil }

type NullFilter struct { }
func (f *NullFilter) Run(raw *bytes.Buffer, filtered *Filtered) {
  filtered.ReadFrom(raw)
}

type Filterable interface {
  Run(raw *bytes.Buffer, filtered *Filtered)
}

type Filter struct {
  runner    string
  script    string
}

func (f *Filter) getCommand(raw *bytes.Buffer) (cmd *exec.Cmd) {
  filename := f.writeScriptFile()

  commands := strings.Fields(f.runner)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd = exec.Command(commands[0], args...)
  cmd.Stdin = raw

  return
}

func (f *Filter) runCommand(cmd *exec.Cmd) (scripted []byte){
  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(string(scripted)); panic(err) }

  return scripted
}

func (f *Filter) writeResult(scripted []byte, filtered *Filtered) {
  noTrailing := strings.TrimSuffix(string(scripted), NL)
  fmt.Fprint(filtered, noTrailing)
}

func (f *Filter) writeScriptFile() string {
  tmp, err := ioutil.TempFile("", PREFIX)
  if err != nil { panic(err) }

  defer tmp.Close()
  _, err  = tmp.WriteString(f.script)
  if err != nil { panic(err) }

  return tmp.Name()
}
