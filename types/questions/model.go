package questions

//Questions question list
type Questions []Question

//Question structure
type Question struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Answer string `json:"answer"`
}

func (qs Questions) GetRandomQuestion() (*Question, error) {
	return nil, nil
}
