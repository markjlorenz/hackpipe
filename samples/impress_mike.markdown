## Write from vim to campfire
```vim
:w !hackpipe -a=campfire
```

## Get images in the browser
https://github.com/prasmussen/chrome-cli

```bash
hackpipe -a=campfire -o="puts JSON.parse(ARGF.read)['body']" | grep --line-buffered -i '^http' | xargs -L1 chrome-cli open
```

## Desktop notifications when I'm mentioned
brew install terminal-notifier

```bash
hackpipe -a=campfire -o="puts JSON.parse(ARGF.read)['body']" | grep --line-buffered -i 'mark' | while read line; do terminal-notifier -message "$line"; done
```

## In combination with jq
[jq](http://stedolan.github.io/jq/manual/#Invokingjq) is a great program for working with JSON data.  I frequently use it to pretty-print with its `.` filter

In your `~/.hackpiperc`
```
  output:
    runner: "jq -C -f "
    script: "."
```
