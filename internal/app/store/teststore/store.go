package teststore

import (
	"github.com/IFtech-A/http_rest_api/internal/app/model"
	"github.com/IFtech-A/http_rest_api/internal/app/store"
	_ "github.com/lib/pq" // something
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
