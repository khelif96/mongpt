package userInput

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func PromptForDatabase(databases []string) string {
	chosenDB := ""

	prompt := &survey.Select{
		Message: "Select Database",
		Options: databases,
		Default: databases[0],
	}
	for chosenDB == "" {
		err := survey.AskOne(prompt, &chosenDB)
		if err != nil {
			fmt.Printf("Prompt failed %v", err)
			return ""
		}
	}
	return chosenDB
}

func PromptForCollectionsToSample(collections []string) []string {
	// Request the user to select if they want to sample all collections or just a few
	shouldSampleAll := false
	sampleConfirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Would you like to sample all %d collections?", len(collections)),
		Default: true,
	}
	err := survey.AskOne(sampleConfirmPrompt, &shouldSampleAll)
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		return nil
	}
	if shouldSampleAll {
		return collections
	}
	// If the user wants to sample some collections, prompt them to select which ones
	collectionsToSample := []string{}
	collectionPrompt := &survey.MultiSelect{
		Message: "Select Collections to Sample",
		Options: collections,
	}
	for len(collectionsToSample) == 0 {
		err = survey.AskOne(collectionPrompt, &collectionsToSample)
		if err != nil {
			fmt.Printf("Prompt failed %v", err)
			return nil
		}
		if len(collectionsToSample) == 0 {
			fmt.Println("You must select at least one collection to sample!")
		}
	}
	return collectionsToSample
}

// PromptForCollectionToSample prompts the user to select a collection to sample
// This only returns one collection in a slice to make it convenient to swap out with
// PromptForCollectionsToSample
func PromptForCollectionToSample(collections []string) []string {
	collectionToSample := ""
	collectionPrompt := &survey.Select{
		Message: "Select Collection to Sample",
		Options: collections,
	}
	for len(collectionToSample) == 0 {
		err := survey.AskOne(collectionPrompt, &collectionToSample)
		if err != nil {
			fmt.Printf("Prompt failed %v", err)
			return nil
		}
		if len(collectionToSample) == 0 {
			fmt.Println("You must select a collection to sample!")
		}
	}
	return []string{collectionToSample}
}
func PromptForAllowingExpensiveQueries(errorMessage string) bool {
	allowExpensiveQueries := false
	expensiveQueryPrompt := &survey.Confirm{
		Message: errorMessage,
		Default: false,
	}
	err := survey.AskOne(expensiveQueryPrompt, &allowExpensiveQueries)
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		return false
	}
	return allowExpensiveQueries
}

func PromptForQuery() string {
	query := ""
	queryPrompt := &survey.Input{
		Message: "What would you like to ask MonGPT:",
	}
	err := survey.AskOne(queryPrompt, &query)
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		return ""
	}
	return query
}
