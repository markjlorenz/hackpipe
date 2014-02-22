package opts

import (
  "flag"
  "os"
  "io/ioutil"
  "encoding/json"
)

type Options struct {
  API     string
  Command string
  Auth    string
  Headers []string
  Input   ioOptions
  Output  ioOptions
}

type ioOptions struct {
  Endpoint string
  Script   string
}

func (o *Options) Parse() {
  // flag.Usage = usage
  var api     string
  var script  string

  flag.StringVar(&api, "a", api, "The API to access")
  flag.StringVar(&script, "e", script, "A script to process the input/output")
  flag.Parse()

  // load in the config file and apply it's settings
  rcFilename   := os.Getenv("HOME")+"/.hackpiperc"
  rcFile, noRc := os.Open(rcFilename)
  if noRc != nil { panic("a `$HOME/.hackpiperc` file is required.") }
  defer rcFile.Close()

  rcBytes, err := ioutil.ReadAll(rcFile)
  if err != nil { panic(err) }

  o = unmarshalConfig(rcBytes)

  // user overrides
  o.API           = api
  o.Input.Script  = script
  o.Output.Script = script
}

func unmarshalConfig(data []byte) (options *Options) {
  err := json.Unmarshal(data, &options)
  if err != nil { panic(err) }
  return
}

// func usage() {
//   fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
//   fmt.Fprintln(os.Stderr, "Takes one position argument, a tweet id")
//   flag.PrintDefaults()
//   fmt.Fprintln(os.Stderr, "Example: ")
//   fmt.Fprintln(os.Stderr, "  ttyttr 434135040256008192")
//   fmt.Fprintln(os.Stderr, "  requests a tweet from: https://twitter.com/someuser/status/434135040256008192")
// }
