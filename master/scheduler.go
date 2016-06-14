package master

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/richardlt/hackathon/types"
	"github.com/richardlt/hackathon/types/questions"
)

func AskQuestions() {
	err := qs.LoadFromFile("questions.json")
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(time.Second * 2)
	for _ = range ticker.C {
		question, _ := qs.GetRandomQuestion()
		fmt.Printf("Question : %v - Réponse : %v\n", question.Title, question.Answer)
		for _, client := range clients {
			go requestClient(client, question)
		}
	}
}

func requestClient(client *types.Client, question questions.Question) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(question)
	res, err := http.Post(client.URL+"/answer", "application/json; charset=utf-8", b)
	if err != nil {
		client.Score = client.Score - 5
	} else {
		var answer types.Answer
		defer res.Body.Close()
		jsonParser := json.NewDecoder(res.Body)

		if err = jsonParser.Decode(&answer); err == nil {
			fmt.Println(answer.Value)
			if answer.Value == question.Answer {
				client.Score = client.Score + 10
				return
			}
		} else {
			fmt.Println(err)
		}
		client.Score = client.Score - 2
	}
	for _, v := range clients {
		fmt.Println(v.URL + " : " + strconv.Itoa(v.Score))
	}
	return
}
