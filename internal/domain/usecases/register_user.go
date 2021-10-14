package usecases

type RegisterUser struct {
}

func (ru RegisterUser) Execute(userId string) error {
	return nil
}

func NewRegisterUser() RegisterUser {
	return RegisterUser{}
}
