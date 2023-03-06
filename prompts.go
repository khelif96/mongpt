package main

import (
	"fmt"

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
	// promptui
	return []string{}
}
