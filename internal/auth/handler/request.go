package handler

type LoginBody struct {
	Nip      string `json:"nip" example:"ARN-2402001"`
	Password string `json:"password" example:"123"`
}

type RegisterBody struct {
	Nip      string `json:"nip" example:"ARN-2402001"`
	Password string `json:"password" example:"123"`
	Name     string `json:"name" example:"Teguh"`
}
