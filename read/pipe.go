package read

import (
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/filter"
  "fmt"
)

type Opts struct {
  Command   string
  InScript  string
  OutScript string
}

func Pipe(network *api.Output, opts *Opts) *filter.Filtered {
  for {
    line, err := network.ReadString('\r')
    if err != nil { panic(err) }
    fmt.Println(line)
  }
}
