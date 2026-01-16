package regmodel

type RegInput struct {
	FirstName  string
	MiddleName string
	LastName   string
	Password   string
	Email      string
}

func NewRegInput(firstName, middleName, lastName, password, email string) (*RegInput, error) {
	if err := validateData(firstName, lastName, password, email); err != nil {
		return nil, err
	}

	return &RegInput{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   password,
		Email:      email,
	}, nil
}

func validateData(firstName, lastName, password, email string) error {
	if firstName == "" {
		return ErrEmptyFirstName
	} else if lastName == "" {
		return ErrEmptyLastName
	} else if password == "" {
		return ErrEmptyPassword
	} else if email == "" {
		return ErrEmptyEmail
	}
	return nil
}
