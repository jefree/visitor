package main

import (
  "github.com/mitchellh/go-mruby"
  "runtime"
  "path/filepath"
  "io/ioutil"
  "log"
)

var (
  _, b, _, _ = runtime.Caller(0)
  basepath = filepath.Dir(b)
)

func launch(scriptPath string) {
  mrb := mruby.NewMrb()

  defer mrb.Close()
  defer quitWebDrivers()

  initVisitorClass(mrb)
  loadScriptFile(mrb, scriptPath)
}

func quitWebDrivers() {
  for _, visitor := range visitors {
    (*visitor.webDriver).Quit()
  }
}

func loadScriptFile(mrb *mruby.Mrb, scriptPath string) {
  fileContent, err := ioutil.ReadFile(scriptPath)
  if err != nil {
    log.Fatal(err)
  }

  fileString := string(fileContent)
  _, err = mrb.LoadString(fileString)

  if err != nil {
    log.Fatal(err)
  }
}
