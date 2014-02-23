package read

import (
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/filter"
  "bytes"
)

const RL byte = '\r'

type Opts struct {
  Runner    string
  OutScript string
}

type Callback func(...interface {}) (int, error) // matches fmt.Print

func Pipe(network *api.Output, opts *Opts, cb Callback) {
  outFilter := filter.NewOutputFilter(opts.Runner, opts.OutScript)

  for {
    line, err := network.ReadString(RL)
    if err != nil { panic(err) }

    lineBuffer := bytes.NewBuffer([]byte(line))
    filtered   := new(filter.Filtered)
    outFilter.Run(lineBuffer, filtered)
    cb( filtered.String() )
  }
}
