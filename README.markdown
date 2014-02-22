# Hack|Pipe

## Examples

Read from the campfire stream:
```
hackpipe -a=c
```

Post a message to campfire:
```
echo "Hi Mark!" | hackpipe -a=c
```

Post a Gist to github:
```
hackpipe < test/github.json -a=g
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

Override your .hackpiperc:
```
hackpipe -a=c -r='node' -o='
process.stdin.on("readable", function(chunk){
    chunk = process.stdin.read();
    if(!chunk){ return };
    console.log(chunk.toString().toUpperCase())
  })
'
```
