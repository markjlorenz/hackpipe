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
  Host    string
  Yolo    bool
  Headers []string
  Input   ioOptions
  Output  ioOptions
}

type ApiOptions map[string]ApiOption
type ApiOption map[string]Options

type ioOptions struct {
  Host    string // overrides the top level host
  Path    string
  Script  string
}

func Parse() (o *Options){
  // flag.Usage = usage
  var api       string
  var inScript  string
  var outScript string

  flag.StringVar(&api, "a", api, "The API to access")
  flag.StringVar(&inScript, "i", inScript, "A script to process the input/output")
  flag.StringVar(&outScript, "o", outScript, "A script to process the input/output")
  flag.Parse()

  assertRequired(api)

  rcBytes  := loadJSON()
  fullOpts := unmarshalConfig(rcBytes)

  op := (*fullOpts)["apis"][api]
  o   = &op

  if o.Input.Host  == "" { o.Input.Host  = o.Host }
  if o.Output.Host == "" { o.Output.Host = o.Host }

  // user overrides
  if inScript  != "" { o.Input.Script  = inScript }
  if outScript != "" { o.Output.Script = outScript }

  return
}

func unmarshalConfig(data []byte) (options *ApiOptions) {
  err := json.Unmarshal(data, &options)
  if err != nil { panic(err) }
  return
}

func loadJSON() []byte {
  rcFilename   := os.Getenv("HOME")+"/.hackpiperc"
  rcFile, noRc := os.Open(rcFilename)
  if noRc != nil { panic("a `$HOME/.hackpiperc` file is required.") }
  defer rcFile.Close()

  rcBytes, err := ioutil.ReadAll(rcFile)
  if err != nil { panic(err) }

  return rcBytes
}

func assertRequired(reqd string) {
  if reqd == "" {
    println("The -a option is requried")
    os.Exit(1)
  }
}

// func usage() {
//   fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
//   fmt.Fprintln(os.Stderr, "Takes one position argument, a tweet id")
//   flag.PrintDefaults()
//   fmt.Fprintln(os.Stderr, "Example: ")
//   fmt.Fprintln(os.Stderr, "  ttyttr 434135040256008192")
//   fmt.Fprintln(os.Stderr, "  requests a tweet from: https://twitter.com/someuser/status/434135040256008192")
// }
