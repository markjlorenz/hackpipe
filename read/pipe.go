package read

import (
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/filter"
  "bytes"
  "fmt"
)

type Opts struct {
  Command   string
  OutScript string
}

func Pipe(network *api.Output, opts *Opts) {
  outFilter := filter.NewFilter(opts.Command, opts.OutScript)
  filtered  := new(filter.Filtered)

  for {
    line, err := network.ReadString('\r')
    if err != nil { panic(err) }

    lineBuffer := bytes.NewBuffer([]byte(line))
    outFilter.Filter(lineBuffer, filtered)
    fmt.Print(filtered)
  }
}
