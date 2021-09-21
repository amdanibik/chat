package users

type UserRegister struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password []byte `form:"password"`
}
