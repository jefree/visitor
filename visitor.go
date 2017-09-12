package main

import (
  "github.com/mitchellh/go-mruby"
  "fmt"
  "log"
  "strings"
  "io/ioutil"
  "sourcegraph.com/sourcegraph/go-selenium"
)

type Visitor struct {
  webDriver *selenium.WebDriver
}

var _id = 0
var visitors = make(map[int]Visitor)

/*
  Initialization Visitor functions
*/

func initVisitorClass(mrb *mruby.Mrb) {
  loadVisitorClassFile(mrb)

  visitorClass := mrb.Class("Visitor", nil)

  visitorClass.DefineClassMethod("create", visitorClassCreate, mruby.ArgsReq(0))
  visitorClass.DefineMethod("visit", visitorVisit, mruby.ArgsReq(0))
}

func loadVisitorClassFile(mrb *mruby.Mrb) {
  fileContent, err := ioutil.ReadFile(strings.Join([]string{basepath, "/visitor.rb"}, ""))
  if err != nil {
    log.Fatal(err)
  }

  fileString := string(fileContent)
  _, err = mrb.LoadString(fileString)
  if err != nil {
    log.Fatal(err)
  }
}

/*
  Visitor API methods
*/

func visitorClassCreate(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
  _id++
  instance, _ := self.Call("new", mruby.Int(_id))

  webDriver := createWebDriver()
  visitors[_id] = Visitor{webDriver}

  return instance, nil
}

func visitorVisit(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
  _id, err := self.Call("id")

  if err != nil {
    log.Fatal(err)
  }

  args := m.GetArgs()

  id := _id.Fixnum()
  var webDriver = visitors[id].webDriver

  url := args[0].String()
  (*webDriver).Get(url)

  title, err := (*webDriver).Title()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(title)

  return nil, nil
}

