// Copyright (C) 2013-2017, The MetaCurrency Project (Eric Harris-Braun, Arthur Brock, et. al.)
// Use of this source code is governed by GPLv3 found in the LICENSE file
//---------------------------------------------------------------------------------------
// command line interface to developing and testing holochain applications

package main

import (
  // "errors"
  fmt         "fmt"
  // holo "github.com/metacurrency/holochain"
  // "github.com/metacurrency/holochain/ui"
  cli         "github.com/urfave/cli"
  os          "os"
  // "os/user"
  // exec        "os/exec"
  syscall     "syscall"
  // "path"
  filepath    "path/filepath"
  // "time"
)

const (
  defaultPort = "4141"
)

// ## forget this nonsense, we only need to get "$GOPATH"
// func LocationOfExecutable() (string) {
//     dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
//     return dir
// }
// func LocationOfBinScripts() (string) {
//     dir := LocationOfExecutable()
//     dir = filepath.Join(dir, "../../bin")
//     return dir
// }
func getCurrentDirectory() (string) {
  dir, err := os.Getwd()
  if err != nil {
    fmt.Println("HC: could not find current directory. Weird. Exitting")
    os.Exit(1)
  }

  return dir
}

func syscallExec(binaryFile string, args ...string){
  syscall.Exec(binaryFile, append([]string{binaryFile}, args...), os.Environ())
}

var debug bool
var rootPath, devPath, name string
var coreBranchInstallPath, appTestScenario string

func setupApp() (app *cli.App) {
  app = cli.NewApp()
  app.Name = "hcdev"
  app.Usage = "holochain dev command line tool"
  // app.Version = fmt.Sprintf("0.0.0 (holochain %s)", holo.VersionStr)

  app.Commands = []cli.Command{
    {
      Name:      "core",
      Usage:     fmt.Sprintf("commands relating to Holochain Core"),
      Subcommands: []cli.Command{
        {
          Name:               "branch",
          // Aliases:    []string{"branch"},
          Usage:              "install a Holochain Core from a git branch",
          Flags:              []cli.Flag{
              cli.StringFlag{
                Name:          "sourceDirectory",
                Usage:         "path to source files containing checked out git branch, defaults to current directory",
                Value:         getCurrentDirectory(),
                Destination:   &coreBranchInstallPath,
              },
          },
          Subcommands:    []cli.Command{
            {
              Name:           "install",
              // Aliases:    []string{"branch"},
              Usage:          "install the version of Holochain Core in '.' onto the host system",
              Action:         func(c *cli.Context) error {
                  fmt.Printf  ("HC: core.branch.install: installing from: %v\n",      coreBranchInstallPath)

                  err := os.Chdir(coreBranchInstallPath)
                  if err != nil {
                    fmt.Printf("HC: core.branch.install: could not change dir to %v", coreBranchInstallPath)
                    os.Exit(1)
                  }

                  // terminates go process
                  syscallExec(
                      filepath.Join(
                        os.Getenv("GOPATH"), 
                        "src/github.com/metacurrency/holochain", 
                        "bin", 
                        "holochain.core.testing.branch",
                      ),
                  )

                  return nil
              },
            },
          },
        },
      },
    },
    {
      Name:      "app",
      Usage:     fmt.Sprintf("commands relating to Holochain Core"),
      Subcommands: []cli.Command{
        {
          Name:               "init",
          // Aliases:    []string{"branch"},
          Usage:              "initialise a directory to be a Holochain App",
          Action:         func(c *cli.Context) error {
              // terminates go process
              syscallExec(
                  filepath.Join(
                    os.Getenv("GOPATH"), 
                    "src/github.com/metacurrency/holochain", 
                    "bin", 
                    "holochain.app.init",
                  ),
                  appTestScenario,
              )

              return nil
          },
        },
        {
          Name:               "testScenario",
          // Aliases:    []string{"branch"},
          Usage:              "run a scenario test",
          Flags:              []cli.Flag{
              cli.StringFlag{
                Name:          "scenario",
                Usage:         "the name of the directory containing the scenario",
                Value:         "",
                Destination:   &appTestScenario,
              },
          },
          Action:         func(c *cli.Context) error {
              // terminates go process
              syscallExec(
                  filepath.Join(
                    os.Getenv("GOPATH"), 
                    "src/github.com/metacurrency/holochain", 
                    "bin", 
                    "holochain.app.testScenario",
                  ),
                  appTestScenario,
              )

              return nil
          },
        },
      },
    },
  }

  return app
}

// func setupApp() (app *cli.App) {
//  app = cli.NewApp()
//  app.Name = "hcdev"
//  app.Usage = "holochain dev command line tool"
//  app.Version = fmt.Sprintf("0.0.0 (holochain %s)", holo.VersionStr)

//  var service *holo.Service

//  app.Flags = []cli.Flag{
//    cli.BoolFlag{
//      Name:        "debug",
//      Usage:       "debugging output",
//      Destination: &debug,
//    },
//    cli.StringFlag{
//      Name:        "execpath",
//      Usage:       "path to holochain dev execution directory (default: ~/.holochaindev)",
//      Destination: &rootPath,
//    },
//    cli.StringFlag{
//      Name:        "path",
//      Usage:       "path to chain source definition directory (default: current working dir)",
//      Destination: &devPath,
//    },
//  }
//  app.Commands = []cli.Command{
//    {
//      Name:      "test",
//      Aliases:   []string{"t"},
//      ArgsUsage: "no args run's all stand-alone | [test file prefix] | [scenario] [role]",
//      Usage:     "run chain's stand-alone or scenario tests",
//      Action: func(c *cli.Context) error {
//        var err error
//        var h *holo.Holochain
//        h, err = getHolochain(c, service)

//        args := c.Args()
//        var errs []error

//        if len(args) == 2 {
//          dir := h.TestPath() + "/" + args[0]
//          role := args[1]

//          err, errs = h.TestScenario(dir, role)
//          if err != nil {
//            return err
//          }
//        } else if len(args) == 1 {
//          errs = h.TestOne(args[0])
//        } else if len(args) == 0 {
//          errs = h.Test()
//        } else {
//          return errors.New("test: expected 0 args (run all stand-alone tests), 1 arg (a single stand-alone test) or 2 args (scenario and role)")
//        }

//        var s string
//        for _, e := range errs {
//          s += e.Error()
//        }
//        if s != "" {
//          return errors.New(s)
//        }
//        return nil
//      },
//    },
//    {
//      Name:      "web",
//      Aliases:   []string{"serve", "w"},
//      ArgsUsage: "[port]",
//      Usage:     fmt.Sprintf("serve a chain to the web on localhost:<port> (defaults to %s)", defaultPort),
//      Action: func(c *cli.Context) error {
//        h, err := getHolochain(c, service)
//        if err != nil {
//          return err
//        }
//        h, err = service.GenChain(name)
//        if err != nil {
//          return err
//        }

//        var port string
//        if len(c.Args()) == 0 {
//          port = defaultPort
//        } else {
//          port = c.Args()[0]
//        }
//        fmt.Printf("Serving holochain with DNA hash:%v on port:%s\n", h.DNAHash(), port)

//        err = h.Activate()
//        if err != nil {
//          return err
//        }
//        //        go h.DHT().HandleChangeReqs()
//        go h.DHT().HandleGossipWiths()
//        go h.DHT().Gossip(2 * time.Second)
//        ui.NewWebServer(h, port).Start()
//        return err
//      },

//       {
//       Name:      "core",
//       Usage:     fmt.Sprintf("commands relating to Holochain Core"),
//       Subcommands: []cli.Command{
//         {
//           Name:       "localSource",
//           // Aliases:    []string{"branch"},
//           Usage:      "Holochain Core commands relating to localSource code",
//           Subcommands: []cli.Command{
//             {
//               Name:       "install",
//               // Aliases:    []string{"branch"},
//               Usage:      "install the version of Holochain Core in '.' onto the host system",
//               Action:     func(c *cli.Context) error {
//                 fmt.Println("calling Script")
                


//                 return nil
//               },
//             },
//           },
//         },
//       },
//    }}

//  app.Before = func(c *cli.Context) error {
//    if debug {
//      os.Setenv("DEBUG", "1")
//    }
//    holo.Initialize()

//    var err error
//    if devPath == "" {
//      devPath, err = os.Getwd()
//      if err != nil {
//        return err
//      }
//    }
//    name = path.Base(devPath)
//    // TODO confirm devPath is actually a holochain app directory

//    if rootPath == "" {
//      rootPath = os.Getenv("HOLOPATH")
//      if rootPath == "" {
//        u, err := user.Current()
//        if err != nil {
//          return err
//        }
//        userPath := u.HomeDir
//        rootPath = userPath + "/" + holo.DefaultDirectoryName + "dev"
//      }
//    }
//    if !holo.IsInitialized(rootPath) {
//      service, err = holo.Init(rootPath, holo.AgentName("test@example.com"))
//      if err != nil {
//        return err
//      }
//      fmt.Println("Holochain dev service initialized:")
//      fmt.Printf("    %s directory created\n", rootPath)
//      fmt.Printf("    defaults stored to %s\n", holo.SysFileName)
//      fmt.Println("    key-pair generated")
//      fmt.Printf("    default agent stored to %s\n", holo.AgentFileName)

//    } else {
//      service, err = holo.LoadService(rootPath)
//    }
//    return err
//  }

//  app.Action = func(c *cli.Context) error {
//    cli.ShowAppHelp(c)

//    return nil
//  }
//  return

// }

func main() {
  app := setupApp()

  app.EnableBashCompletion = true

  err := app.Run(os.Args)
  if err != nil {
   fmt.Printf("Error: %v\n", err)
   os.Exit(1)
  }
}

// func getHolochain(c *cli.Context, service *holo.Service) (h *holo.Holochain, err error) {
//  fmt.Printf("Copying chain to: %s\n", rootPath)
//  err = os.RemoveAll(rootPath + "/" + name)
//  if err != nil {
//    return
//  }
//  h, err = service.Clone(devPath, rootPath+"/"+name, false)
//  if err != nil {
//    return
//  }
//  h, err = service.Load(name)
//  if err != nil {
//    return
//  }
//  return
// }
