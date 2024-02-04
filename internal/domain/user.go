package domain

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
}

type CreateUserOutput struct {
}
