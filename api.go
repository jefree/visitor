package visitor

import (
  "github.com/mitchellh/go-mruby"
  "runtime"
  "path/filepath"
)

var (
  _, b, _, _ = runtime.Caller(0)
  basepath = filepath.Dir(b)
)

func Init(mrb *mruby.Mrb) {
  initVisitorClass(mrb)

  for _, visitor := range visitors {
    (*visitor.webDriver).Quit()
  }
}
