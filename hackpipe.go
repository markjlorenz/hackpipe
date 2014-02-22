package main

import (
  "dapplebeforedawn/hackpipe/write"
  "dapplebeforedawn/hackpipe/read"
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/options"
  "fmt"
)

func main() {
  opts := opts.Parse()

  go reader(opts)
  writer(opts)
}

func writer(opts *opts.Options) {
  writeOpts := &api.Opts{
    Path:    opts.Input.Path,
    Auth:    opts.Auth,
    Headers: opts.Headers,
    Host:    opts.Input.Host,
    Yolo:    opts.Yolo,
  }
  input := api.NewInput(writeOpts)

  inputOpts := &write.Opts{
    Runner:    opts.Runner,
    InScript:  opts.Input.Script,
    OutScript: opts.Output.Script,
  }
  afterWrite := write.Pipe(input, inputOpts)

  fmt.Println(afterWrite)
}

func reader(opts *opts.Options) {
  // if no streaming endpoint configured, don't do anthing.
  if opts.Output.Path == "" { return }

  readOpts := &api.Opts{
    Path:    opts.Output.Path,
    Auth:    opts.Auth,
    Headers: opts.Headers,
    Host:    opts.Output.Host,
    Yolo:    opts.Yolo,
  }
  readable := api.NewOutput(readOpts)

  outputOpts := &read.Opts{
    Runner:    opts.Runner,
    OutScript: opts.Output.Script,
  }

  read.Pipe(readable, outputOpts, fmt.Println)
}
