package service

type AuthService interface {
	Login(username, password string) (string, error)
}