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
    Scheme:  opts.Input.Scheme,
    Path:    opts.Input.Path,
    Auth:    opts.Auth,
    Headers: opts.Headers,
    Query:   opts.Input.Query,
    Method:  opts.Input.Method,
    Host:    opts.Input.Host,
    Yolo:    opts.Yolo,
  }
  input := api.NewInput(writeOpts)

  inputOpts := &write.Opts{
    InRunner:  opts.Input.Runner,
    OutRunner: opts.Output.Runner,
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
    Scheme:  opts.Output.Scheme,
    Path:    opts.Output.Path,
    Auth:    opts.Auth,
    Headers: opts.Headers,
    Host:    opts.Output.Host,
    Method:  opts.Output.Method,
    Yolo:    opts.Yolo,
  }
  readable := api.NewOutput(readOpts)

  outputOpts := &read.Opts{
    Runner:    opts.Output.Runner,
    OutScript: opts.Output.Script,
  }

  read.Pipe(readable, outputOpts, fmt.Println)
}
