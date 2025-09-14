package handler

type GetPracticesQuery struct {
	Limit  string `json:"limit" example:"10"`
	Offset string `json:"offset" example:"1"`
}
