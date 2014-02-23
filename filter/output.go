package filter

import (
  "bytes"
)

type OutputFilter struct {
  Filter
}

func NewOutputFilter(runner, script string) Filterable {
  // if we can't filter, don't try.
  if runner == "" || script == "" {
    return &NullFilter{}
  }

  return &OutputFilter {
    Filter{
      runner: runner,
      script: script,
    },
  }
}

func (f *OutputFilter) Run(raw *bytes.Buffer, filtered *Filtered) {
  cmd      := f.getCommand(raw)
  scripted := f.runCommand(cmd)
  f.writeResult(scripted, filtered)
}
