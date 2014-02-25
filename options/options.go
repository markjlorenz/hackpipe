package opts

import (
  "flag"
  "os"
  "io/ioutil"
  "launchpad.net/goyaml"
)

type ApiOptions map[string]ApiOption
type ApiOption map[string]Options

type Options struct {
  API     string
  Runner  string
  Auth    string
  Host    string
  Yolo    bool
  Headers map[string]string
  Input   ioOptions
  Output  ioOptions
}

type ioOptions struct {
  Host    string // overrides the top level host
  Runner  string // overrides the top level runner
  Path    string
  Query   string
  Method  string
  Script  string
}

func Parse() (o *Options){
  // flag.Usage = usage
  var api           string
  var inScript      string
  var outScript     string
  var runner        string
  var inputRunner   string
  var outputRunner  string

  flag.StringVar(&api, "a", api, "The API to access")
  flag.StringVar(&inScript, "i", inScript, "A script to pre-process the api input")
  flag.StringVar(&outScript, "o", outScript, "A script to process the api output")
  flag.StringVar(&runner, "r", runner, "The program that runs your scripts.  The data will be availble on STDIN.")
  flag.StringVar(&inputRunner, "ri", inputRunner, "The same as '-r', but only applied to input")
  flag.StringVar(&outputRunner, "ro", outputRunner, "The same as '-r', but only applied to output")
  flag.Parse()

  assertRequired(api)

  rcBytes  := loadJSON()
  fullOpts := unmarshalConfig(rcBytes)

  op := (*fullOpts)["apis"][api]
  o   = &op

  // defaults
  if o.Input.Method == "" { o.Input.Method = "POST" }

  // user overrides
  if inScript     != "" { o.Input.Script  = inScript }
  if outScript    != "" { o.Output.Script = outScript }
  if runner       != "" { o.Runner        = runner }
  if inputRunner  != "" { o.Input.Runner  = inputRunner }
  if outputRunner != "" { o.Output.Runner = outputRunner }

  // shadowing
  if o.Input.Host    == "" { o.Input.Host  = o.Host }
  if o.Output.Host   == "" { o.Output.Host = o.Host }
  if o.Input.Runner  == "" { o.Input.Runner  = o.Runner }
  if o.Output.Runner == "" { o.Output.Runner = o.Runner }

  return
}

func unmarshalConfig(data []byte) (options *ApiOptions) {
  err := goyaml.Unmarshal(data, &options)
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
