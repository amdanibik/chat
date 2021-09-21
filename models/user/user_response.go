package user

type UserResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
