package userInput

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/khelif96/mongpt/db"
)

func PromptForDatabase(databases []string) string {
	chosenDB := ""

	prompt := &survey.Select{
		Message: "Select Database",
		Options: databases,
		Default: databases[0],
	}
	err := survey.AskOne(prompt, &chosenDB)
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		return ""
	}
	fmt.Println(chosenDB)

	err = db.ChooseDatabase(chosenDB)
	if err != nil {
		fmt.Println(err.Error())
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
