package store

type User struct {
	Id    string
	Email string
}

type Store interface {
	Users() ([]User, error)
	User(id string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
}
