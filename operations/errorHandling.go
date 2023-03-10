package operations

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/khelif96/mongpt/gpt"
	"gopkg.in/mgo.v2/bson"
)

func CleanGPTResponse(response string) ([]bson.M, error) {
	finalResponse := []bson.M{}
	backtickPattern := regexp.MustCompile("(?s)```(.+?)```")
	// Use the regexp package to search for the pattern in the string
	match := backtickPattern.FindStringSubmatch(response)
	if len(match) > 1 {
		fmt.Println("Found backticks")
		response = match[1]
		// Sometimes chatgpt returns an array but without the brackets. This fixes that.
		arrayPattern := regexp.MustCompile(`(?s)\[(.+?)\]`)
		match := arrayPattern.FindStringSubmatch(response)
		if len(match) == 0 {
			response = fmt.Sprintf("[%s]", response)
		}

		err := json.Unmarshal([]byte(response), &finalResponse)
		if err != nil {
			return nil, err
		}
	} else {
		gptResponse := gpt.ChatGPTResponse{}
		err := json.Unmarshal([]byte(response), &gptResponse)
		if err != nil {
			return nil, err
		}
		finalResponse = gptResponse.Query
	}
	fmt.Println(finalResponse)
	if (len(finalResponse)) == 0 {
		return nil, fmt.Errorf("No response found")
	}
	return finalResponse, nil
}
