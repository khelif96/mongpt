package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evergreen-ci/utility"
	"github.com/urfave/cli/v2" // imports as package "cli"

	"github.com/khelif96/mongpt/db"
	"github.com/khelif96/mongpt/gpt"
	"github.com/khelif96/mongpt/operations"
	"github.com/khelif96/mongpt/userInput"
)

const cacheDirectory = "./bin/cache/"

func main() {
	operations.Init()
	operations.LoadVariablesIntoEnvironment()
	db.Init()
	gpt.Init()
	app := &cli.App{
		Name:  "MonGPT CLI",
		Usage: "A simple CLI program to perform GPT style queries on MongoDB",
		Action: func(c *cli.Context) error {
			return nil
		},
		DefaultCommand: "default",
		Commands: []*cli.Command{
			{
				Name:  "default",
				Usage: "Initial entry point and command",
				Action: func(c *cli.Context) error {
					fmt.Println("Welcome to MonGPT CLI!")
					fmt.Println("Please choose a database to work with:")
					databases, err := db.GetDatabases()
					if err != nil {
						log.Fatal(err)
					}

					database := userInput.PromptForDatabase(databases)
					// operations.ClearTerminal()
					fmt.Println("You chose database: ", database)
					collections := db.GetCollections()
					selectedCollections := userInput.PromptForCollectionsToSample(collections)
					fmt.Println("You chose to sample the following collections: ", selectedCollections)

					for _, collection := range selectedCollections {
						fmt.Println("Sampling collection: ", collection)
						sample, err := db.CollectDocumentSamplesFromCollection(collection)
						if err != nil {
							log.Fatal(err)
						}
						schema := operations.GetSchemaFromDocument(sample)
						fileName := cacheDirectory + collection + ".json"
						jsonSchema := operations.ConvertBSONToJSON(schema)
						err = operations.WriteJSONSchemaToFile(fileName, jsonSchema)
						if err != nil {
							log.Fatal(err)
						}
					}
					schemas := operations.ReadJSONSchemasFromDir(cacheDirectory)
					if (len(schemas)) == 0 {
						log.Fatal("No schemas found in cache directory")
					}

					query := userInput.PromptForQuery()

					response, err := gpt.AskGPT(operations.FormatSchemas(schemas), query)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("Response: ", utility.FromStringPtr(response))
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
