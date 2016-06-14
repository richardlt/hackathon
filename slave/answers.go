package slave

import (
	"fmt"

	"github.com/richardlt/hackathon/types"
)

var answers map[string]string

func init() {
	answers = make(map[string]string)
	answers["Que donne 1 + 1 ?"] = "2"
}

// Answer answers the question
func Answer(question types.Question) (string, error) {

	val, ok := answers[question.Title]
	if !ok {
		return "", fmt.Errorf("Can't find a answer for the question : %v", question)
	}

	return val, nil
}
