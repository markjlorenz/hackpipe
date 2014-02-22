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

  writeOpts := &api.Opts{
    Path: opts.Input.Path,
    Auth: opts.Auth,
    Host: opts.Host,
    Yolo: opts.Yolo,
  }
  input := api.NewInput(writeOpts)

  inputOpts := &write.Opts{
    Command:   opts.Command,
    InScript:  opts.Input.Script,
    OutScript: opts.Output.Script,
  }
  afterWrite := write.Pipe(input, inputOpts)

  fmt.Print(afterWrite)

//-----------------------------------------------
  fmt.Println("READY")
  readOpts := &api.Opts{
    Path: opts.Output.Path,
    Auth: opts.Auth,
    Host: opts.Host,
    Yolo: opts.Yolo,
  }
  readable := api.NewOutput(readOpts)

  outputOpts := &read.Opts{
    Command:   opts.Command,
    InScript:  opts.Input.Script,
    OutScript: opts.Output.Script,
  }
  afterRead := read.Pipe(readable, outputOpts)

  fmt.Print(afterRead)

}

// if stdin then we are in write mode:
  // pass each line through the input script
  // POST to the endpoint

// when in read mode:
  // connect to the endpoint
  // start getting data
  // pass each ??? through the output script

// CONFIG FILE:
  // - api abbreviation flag
  // - authorization header value
  // - any other header k/v pairs
  // - script runner command
  //
  // - api input:
  //   - endpoint
  //   - script (line oriented) [ or use -e option ]
  // - api output:
  //   - endpoint
  //   - script (message oriented) [ or use -e option ]
