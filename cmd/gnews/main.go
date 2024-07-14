package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/leberjs/gnews/internal/cmd"
	"github.com/leberjs/gnews/internal/config"
	"github.com/urfave/cli/v2"

	bolt "go.etcd.io/bbolt"
)

// commands:
//  - add: add entry to db
//  - rm: remove entry from db
//  - ls: list entries in db
//  - read: read all latest news for all entries in db
//      * option - id: read latest news for only specified entry using id
//      * option - name: read latest news for only specified entry using name
//  - config: take options to set certain config values (maybe)
//  - init: initialize with options (maybe)

func main() {
  var cfg config.Config
  cfg, err := config.GetConfig()
  if os.IsNotExist(err) {
    cfg = config.InitConfig()
  }

  dbPath := filepath.Join(config.ConfigDir(), cfg.Database.Name + ".db")
  db, err := bolt.Open(dbPath, 0600, nil)
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()

  db.Update( func(tx *bolt.Tx) error{
    _, err := tx.CreateBucketIfNotExists([]byte("apps"))
    if err != nil {
      log.Fatal(err)
    }

    return nil
  })

  var (
    appId string
    appName string
  )

  app := &cli.App{
    Name: "gnews",
    Usage: "Get latest news from personally maintained list of Steam games",
    Commands: []*cli.Command{
      {
        Name: "add",
        Usage: "Add entry to database",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "app-id",
            Usage: "Steam game App Id (required)",
            Destination: &appId,
            Required: true,
          },
          &cli.StringFlag{
            Name: "name",
            Usage: "Friendly name for game. (Defaults to app-id)",
            Destination: &appName,
          },
        },
        Action: func(ctx *cli.Context) error {
          cmd.Add(db, appId, appName)

          return nil
        },
      },
      {
        Name: "ls",
        Usage: "List all entries in databse",
        Action: func(ctx *cli.Context) error {
          apps := cmd.List(db)
          res, _ := json.MarshalIndent(apps, "", "    ")

          fmt.Println(string(res))
          return nil
        },
      },
      {
        Name: "read",
        Usage: "Read latest news for entries in database. Defaults to all entries.",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "app-id",
          },
          &cli.StringFlag{
            Name: "name",
          },
        },
        Action: func(ctx *cli.Context) error {
          res := cmd.Read(db)
          
          for _, news := range res {
            fmt.Printf("==== News for %s ====\n", news.Name)
            fmt.Println(news.Contents)
            fmt.Println("========")
            fmt.Println()
          }

          return nil
        },
      },
      {
        Name: "rm",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "app-id",
            Usage: "Steam game App Id (required)",
            Destination: &appId,
            Required: true,
          },
        },
        Usage: "Remove entry from database",
        Action: func(ctx *cli.Context) error {
          cmd.Remove(db, appId)
          return nil
        },
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
