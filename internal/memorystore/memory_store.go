package memorystore

import (
	"errors"
	"strconv"

	"github.com/iamhectorsosa/web-server-demo/internal/store"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrFailedSettingEntry = errors.New("failed setting entry")
)

type Store struct {
	users map[string]store.User
}

func New(initialUsers ...store.User) *Store {
	store := &Store{users: map[string]store.User{}}

	for _, user := range initialUsers {
		store.users[user.Id] = user
	}

	return store
}

func (s *Store) Users() ([]store.User, error) {
	users := make([]store.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}

func (s *Store) User(id string) (store.User, error) {
	user, ok := s.users[id]
	if !ok {
		return store.User{}, ErrNotFound
	}
	return user, nil
}

func (s *Store) CreateUser(user store.User) (store.User, error) {
	newId := strconv.Itoa(len(s.users) + 1)
	newUser := store.User{
		Id:    newId,
		Email: user.Email,
	}
	s.users[newId] = newUser
	return newUser, nil
}

func (s *Store) UpdateUser(user store.User) (store.User, error) {
	if _, err := s.User(user.Id); err != nil {
		return store.User{}, err
	}
	updatedUser := store.User{
		Id:    user.Id,
		Email: user.Email,
	}
	s.users[user.Id] = updatedUser
	return updatedUser, nil
}

func (s *Store) DeleteUser(id string) error {
	user, err := s.User(id)
	if err != nil {
		return err
	}
	delete(s.users, user.Id)
	return nil
}
