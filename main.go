package main

import (
	"encoding/json"
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

					chosenDB := userInput.PromptForDatabase(databases)
					err = db.ChooseDatabase(chosenDB)
					if err != nil {
						fmt.Println(err.Error())
					}

					collections := db.GetCollections()
					selectedCollections := userInput.PromptForCollectionToSample(collections)

					for _, collection := range selectedCollections {
						fmt.Println("Sampling collection: ", collection)
						sample, err := db.CollectDocumentSamplesFromCollection(collection)
						if err != nil {
							log.Fatal(err)
						}

						schema := operations.GetSchemaFromDocument(sample)
						fileName := cacheDirectory + collection + ".json"
						jsonSchema := operations.ConvertBSONToJSON(schema)
						fmt.Println(fmt.Sprintf("The document schema is as follows: %s", jsonSchema))
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

					retries := 0
					chatResponse := gpt.ChatGPTResponse{}

					for true {
						response, err := gpt.AskGPT(operations.FormatSchemas(schemas), query)
						if err != nil {
							log.Fatal(err)
						}

						operations.WriteJSONSchemaToFile("./bin/cache/response.json", utility.FromStringPtr(response))

						// Turn the response into a ChatGPTResponse
						err = json.Unmarshal([]byte(utility.FromStringPtr(response)), &chatResponse)
						if err == nil {
							break
						}
						fmt.Println("Failed to parse response from GPT-3, retrying...")
						retries++
						if retries == 3 {
							log.Fatal("Failed to parse response from GPT-3")
						}
					}

					collection, err := db.ChooseCollection(selectedCollections[0])
					if err != nil {
						log.Fatal(err)
					}

					results, err := db.PerformAggregation(collection, chatResponse.Query)
					if err != nil {
						log.Fatal(err)
					}
					response, err := gpt.AskGPTToReadResponse(operations.ConvertBSONArrayToJSON(results))
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
