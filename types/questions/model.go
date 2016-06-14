package questions

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
)

//Questions question list
type Questions []Question

//Question structure
type Question struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Answer string `json:"answer"`
}

func (qs Questions) GetRandomQuestion() (Question, error) {
	size := len(qs)
	if size <= 0 {
		return Question{}, errors.New("No questions found.")
	}
	index := rand.Intn(size)
	return qs[index], nil

}

func (qs *Questions) Load(path string) error {

	configFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening questions file: %v", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&qs); err != nil {
		return fmt.Errorf("Error parsing questions file", err.Error())
	}
	return nil
}
