package write

import (
  "dapplebeforedawn/hackpipe/api"
  "os"
  "os/exec"
  "bufio"
  "bytes"
  "fmt"
  "time"
  "strings"
  "strconv"
)

type Opts struct {
  Command string
  Script  string
}

type Filtered struct {
  bytes.Buffer
}
func (f *Filtered) Close() error { return nil }

func Pipe(network *api.Input, opts *Opts) {
  raw        := new(bytes.Buffer)
  inFiltered := new(Filtered)
  response   := new(bytes.Buffer)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    fmt.Fprintln(raw, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  filterInput(opts.Command, opts.Script, raw, inFiltered)

  res := network.Write(inFiltered)
  _, err := response.ReadFrom(res)
  if err != nil { panic(err) }

  fmt.Println(response)
}

func filterInput(command, script string, raw *bytes.Buffer, filtered *Filtered) {
  filename := writeScriptFile(script)

  commands := strings.Fields(command)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd := exec.Command(commands[0], args...)
  cmd.Stdin = raw

  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(filtered); panic(err) }

  fmt.Fprint(filtered, string(scripted))
  return
}

func writeScriptFile(script string) (filename string) {
  filename = "/tmp/hackpipe:"+strconv.FormatInt(time.Now().UnixNano(), 10)
  file, _ := os.Create(filename)
  defer file.Close()
  _, err := file.WriteString(script)
  if err != nil { panic(err) }

  return
}
