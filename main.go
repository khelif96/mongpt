package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // imports as package "cli"

	"github.com/khelif96/mongpt/db"
	"github.com/khelif96/mongpt/operations"
)

func main() {
	db.Init()

	app := &cli.App{
		Name:  "MonGPT CLI",
		Usage: "A simple CLI program to perform GPT style queries on MongoDB",
		Action: func(c *cli.Context) error {
			log.Println("Hello friend!")
			return nil
		},
		DefaultCommand: "default",
		Commands: []*cli.Command{
			{
				Name:  "default",
				Usage: "Initial entry point and command",
				Action: func(c *cli.Context) error {
					fmt.Println("Hi there!")
					fmt.Println("Welcome to MonGPT CLI!")
					fmt.Println("Please choose a database to work with:")
					databases, err := db.GetDatabases()
					if err != nil {
						log.Fatal(err)
					}

					database := PromptForDatabase(databases)
					operations.ClearTerminal()
					fmt.Println("You chose database: ", database)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
