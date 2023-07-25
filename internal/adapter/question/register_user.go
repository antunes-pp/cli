package question

type RegisterUser struct {
	Name    string
	Email   string
	Squads  []string
	Confirm bool
}

func (r RegisterUser) GetName() string {
	return r.Name
}

func (r RegisterUser) GetEmail() string {
	return r.Email
}

func (r RegisterUser) GetSquads() []string {
	return r.Squads
}
