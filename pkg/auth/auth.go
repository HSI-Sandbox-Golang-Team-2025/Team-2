package auth

// LoginBody represents the payload for the login endpoint (Postman uses nip/password)
type LoginBody struct {
	NIP      string `json:"nip" example:"ARN-2402001"`
	Password string `json:"password" example:"123"`
}

// RegisterBody represents the payload for the register endpoint (Postman uses nip/password/name)
type RegisterBody struct {
	NIP      string `json:"nip" example:"ARN-2402009"`
	Password string `json:"password" example:"123"`
	Name     string `json:"name" example:"Luthfi"`
}

// TokenData holds the JWT token returned after authentication
type TokenData struct {
	Token string `json:"token" example:"eyJhbGciOiJI..."`
}

// SuccessLoginResponse is the standard success response for login
type SuccessLoginResponse struct {
	Message string    `json:"message" example:"Login success!"`
	Data    TokenData `json:"data"`
}

// InvalidLoginResponse is the error response for failed login
type InvalidLoginResponse struct {
	Message string `json:"message" example:"Invalid credentials!"`
	Error   string `json:"error,omitempty" example:"email or password is incorrect"`
}

// SuccessRegisterResponse is the success response for register (same shape as login)
type SuccessRegisterResponse = SuccessLoginResponse

// InvalidRegisterResponse is the error response for register
type InvalidRegisterResponse = InvalidLoginResponse
