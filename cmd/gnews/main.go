package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name: "gnews",
    Usage: "Get latest news from personally maintained list of Steam games",
    Action: func(*cli.Context) error {
      fmt.Println("wip")
      return nil
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
