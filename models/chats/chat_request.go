package chat

type ChatRecord struct {
	Id      uint   `form:"id"`
	Sender  string `form:"sender"`
	Message string `form:"message"`
}
