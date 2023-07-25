package port

type InputRegisterUser interface {
	GetName() string
	GetEmail() string
	GetSquads() []string
}
