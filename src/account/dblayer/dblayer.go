package dblayer

type DBLayer interface {
	Login()
	Logout()
	IsLogIn()
	AddUser()
	UpdateUser()
	DeleteUser()
	FindLoginId()
	FindPassword()
	ChangePassword()
}
