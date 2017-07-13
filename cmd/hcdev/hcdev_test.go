package main

import (
  fmt        "fmt"

	.          "github.com/smartystreets/goconvey/convey"
	_          "github.com/urfave/cli"
	testing    "testing"
)

func TestSetupApp(t *testing.T) {
	app := setupApp()
	Convey("it should create the cli App", t, func() {
		So(app.Name, ShouldEqual, "hcdev")
	})
}

func TestLocationOfExecutable(t *testing.T) {
  Convey("should print the location of the executable", t, func() {
    fmt.Sprintf("locationOfExecutable: %v", locationOfExecutable() )
  })
}