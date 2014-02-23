# Hackpipe
> A unix pipe like interace for REST and streaming HTTP APIs.

## TL;DR
![](https://github.com/dapplebeforedawn/hackpipe/raw/master/samples/demo.gif)

## Installation
 - Put `bin/hackpipe` into a directory that's in your PATH.
 - Copy `samples/.hackpiperc` to `$HOME/.hackpiperc` and modify it for your APIS

## `~/.hackpiperc`
There is a sample `~/.hackpiperc` included in the `sample` directory.  This should give you a really good overview of what configuration options are availbe to you.

## Examples
These examples assume you are using the sample `.hackpiperc`

----
Read from the campfire stream:
```bash
hackpipe -a=campfire
```

Post a message to campfire:
```bash
echo "Hi Mark!" | hackpipe -a=campfire
```

----
Post a Gist to github, and filter the URL only out of the response:
```bash
hackpipe < test/github.json -a=github -o='puts JSON.parse(ARGF.read)["url"]'
```

[test/github.json]
```text
{
  "description": "the description for this gist",
  "public": false,
  "files": {
    "file1.txt": {
      "content": "String file contents"
    }
  }
}
```

----
Commandline args override the .hackpiperc:
```bash
hackpipe -a=campfire -r='node' -o='
process.stdin.on("readable", function(chunk){
    chunk = process.stdin.read();
    if(!chunk){ return };
    console.log(chunk.toString().toUpperCase())
  })
'
```

----
The input script can modify the request by writing to some special files that are accessible through enviroment varables.

Using the query string global variable:
```bash
# the sample .hackpiperc file sets some common query string values for us
# we want to append a few more
hackpipe -a=cmm -e="File.open(ENV['QUERY'], 'a') { |q| q  << '&q=marinol' }"
# press ^D, we don't have any stdin to provide
```

## Notes on Authorization
For the `auth` key in your `~/.hackpiperc` hackpipe expects you to pre-compute any encoding.  For example, many APIs use HTTP Basic auth.  For those APIs your `auth` value would be the result of:
  ```pseudocode
   "Basic " + base64encode("someusername" + ":" + "somepassword")
  ```
In otherwords, the value of `auth` will be set as the value of the `Authorization` header.  This makes it easy for hackpipe to support digest, token and bearer auth as well.
