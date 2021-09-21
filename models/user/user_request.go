package user

type UserRegister struct {
	Id       uint   `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password []byte `form:"password"`
}
