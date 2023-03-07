package userInput

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evergreen-ci/utility"
	"github.com/khelif96/mongpt/db"
	"github.com/manifoldco/promptui"
)

func PromptForDatabase(databases []string) string {
	var chosenDB *string
	for chosenDB == nil {
		prompt := promptui.Select{
			Label: "Select Database",
			Items: databases,
		}

		_, result, err := prompt.Run()
		if result != "" {
			chosenDB = utility.ToStringPtr(result)
		}
		err = db.ChooseDatabase(utility.FromStringPtr(chosenDB))
		if err != nil {
			fmt.Println(err.Error())
			chosenDB = nil
		}
	}
	return utility.FromStringPtr(chosenDB)
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
