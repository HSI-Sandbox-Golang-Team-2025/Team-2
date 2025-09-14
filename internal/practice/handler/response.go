package handler

type SuccessLoginResponse struct {
	Message string `json:"message" example:"Login success!"`
	Data    SuccessLoginData
}

type SuccessLoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUz.."`
}

type InvalidLoginResponse struct {
	Message string `json:"message" example:"Invalid Credentials!"`
}

type SuccessRegisterResponse struct {
	Message string `json:"message" example:"Registration success!"`
	Data    SuccessRegisterData
}

type SuccessRegisterData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUz.."`
}

type InvalidRegisterResponse struct {
	Message string `json:"message" example:"Registration failed!"`
}
