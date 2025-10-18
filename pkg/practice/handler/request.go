package handler

type CreatePracticeBody struct {
	TrackId   uint       `json:"trackId" example:"1"`
	Title     string     `json:"title" example:"Latihan 2"`
	Body      string     `json:"body" example:"Latihan ini anda akan menguji kemampuan React anda"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Question      string         `json:"question" example:"Apakah 1 + 1 = 2"`
	AnswerChoices []AnswerChoice `json:"answerChoices"`
}

type AnswerChoice struct {
	Answer string `json:"answer" example:"Benar"`
}
