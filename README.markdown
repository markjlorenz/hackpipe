# Hack|Pipe
> A unix pipe like interace for REST and streaming HTTP apis.

## `~/.hackpiperc`
There is a sample `~/.hackpiperc` included in the `sample` directory.  This should give you a really good overview of what configuration options are availbe to you.

## Examples
These examples assume you are using the sample `.hackpiperc`

----
Read from the campfire stream:
```
hackpipe -a=c
```

Post a message to campfire:
```
echo "Hi Mark!" | hackpipe -a=c
```

----
Post a Gist to github, and filter the URL only out of the response:
```
hackpipe < test/github.json -a=g -o='puts JSON.parse(ARGF.read)["url"]'
```

[test/github.json]
```
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
```
hackpipe -a=c -r='node' -o='
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
```
# the sample .hackpiperc file sets some common query string values for us
# we want to append a few more
hackpipe -a=cmm -e="File.open(ENV['QUERY'], 'a') { |q| q  << '&q=ambien' }"
# press ^D, we don't have any stdin to provide
```
