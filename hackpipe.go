package main

import (
  "dapplebeforedawn/hackpipe/write"
)

func main() {
  writeOpts := &write.Opts{}
  write.Pipe(writeOpts)
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
