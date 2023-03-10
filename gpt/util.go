package gpt

import (
	"math"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func calculateTokenUsage(tokens ...string) int {
	// 4 characters represent 1 token and tokens are rounded up to the nearest whole number
	return int(math.Ceil(float64(len(strings.Join(tokens, ""))) / 4))
}

func calculateGPT3Cost(tokenCount int) float64 {
	// gpt3 costs $0.002 per 1000 tokens
	return float64(tokenCount) * 0.002 / 1000

}

type ChatGPTResponse struct {
	Query       []bson.M `json:"query"`
	Explanation string   `json:"explanation"`
}
