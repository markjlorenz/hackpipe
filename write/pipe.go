package write

import (
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/filter"
  "os"
  "bufio"
  "bytes"
  "fmt"
)

type Opts struct {
  Runner    string
  InScript  string
  OutScript string
}

func Pipe(network *api.Input, opts *Opts) *filter.Filtered {
  raw         := new(bytes.Buffer)
  inFiltered  := new(filter.Filtered)
  outFiltered := new(filter.Filtered)
  response    := new(bytes.Buffer)
  inFilter    := filter.NewFilter(opts.Runner, opts.InScript)
  outFilter   := filter.NewFilter(opts.Runner, opts.OutScript)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    fmt.Fprintln(raw, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  inFilter.Filter(raw, inFiltered)

  res := network.Write(inFiltered)
  _, err := response.ReadFrom(res)
  if err != nil { panic(err) }

  outFilter.Filter(response, outFiltered)

  return outFiltered
}
