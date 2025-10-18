package handler

type CreateProjectBody struct {
	TrackId uint   `json:"trackId" example:"1"`
	Title   string `json:"title" example:"Final Project"`
	Body    string `json:"body" example:"Anda akan diminta untuk membuat final project"`
}
