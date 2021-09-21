package room

type RoomResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type RoomResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}
