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
}

func Pipe(opts *Opts) {
  command  := "ruby -rjson"
  script   := "j = JSON.parse(ARGF.read); j['description'] = j['description'].upcase; puts j.to_json"
  raw      := new(bytes.Buffer)
  filtered := new(bytes.Buffer)
  response := new(bytes.Buffer)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    fmt.Fprintln(raw, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  filename := "/tmp/hackpipe:"+strconv.FormatInt(time.Now().UnixNano(), 10)
  file, _ := os.Create(filename)
  defer file.Close()
  _, err := file.WriteString(script)
  if err != nil { panic(err) }

  commands := strings.Fields(command)
  args     := append([]string{}, commands[1:]...)
  args      = append(args, filename)

  cmd := exec.Command(commands[0], args...)
  cmd.Stdin = raw

  scripted, err := cmd.CombinedOutput()
  if err != nil { fmt.Println(filtered); panic(err) }

  fmt.Fprint(filtered, string(scripted))

  res := api.NewInput(filtered)
  _, err = response.ReadFrom(res)
  fmt.Println(response)
}
