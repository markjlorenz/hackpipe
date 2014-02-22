package main

import (
  "dapplebeforedawn/hackpipe/write"
  "dapplebeforedawn/hackpipe/read"
  "dapplebeforedawn/hackpipe/api"
  "dapplebeforedawn/hackpipe/options"
  "fmt"
)

func main() {
  opts   := opts.Parse()

  go writer(opts)
  go reader(opts)

  select{ } //chill here for a while
  // for {
  //   select {
  //   case w := <-writes:
  //     println(w)
  //   case r := <-reads:
  //     println(r)
  //   }
  // }
}

func writer(opts *opts.Options) {
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
}

func reader(opts *opts.Options) {
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
