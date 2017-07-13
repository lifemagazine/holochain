package main

import (
  fmt         "fmt"
  strings     "strings"

  .           "github.com/smartystreets/goconvey/convey"
  _           "github.com/urfave/cli"
  testing     "testing"
)

// func TestSetupApp(t *testing.T) {
//  app := setupApp()
//  Convey("it should create the cli App", t, func() {
//    So(app.Name, ShouldEqual, "hcdev")
//  })
// }

func TestLocationOfExecutable(t *testing.T) {
  Convey("should print the location of the executable", t, func() {
    location  := LocationOfExecutable()
    fmt.Sprintf("locationOfExecutable: %v", location)
    fmt.Println(location)
    So( strings.Contains(location, "/cmd/hcdev/") , ShouldBeTrue)
  })

  // ##output
  // /tmp/go-build500264145/_/home/christopher/dev.bin2hcdcev/cmd/hcdev/_test
  // .
  // 1 total assertion

  // PASS
  // ok      _/home/christopher/dev.bin2hcdcev/cmd/hcdev     0.003s

}
