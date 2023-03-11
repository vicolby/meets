package types

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
