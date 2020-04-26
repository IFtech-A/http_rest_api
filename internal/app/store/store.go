package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // something
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		log.Fatal("cant open database")
		return err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cant ping database")
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {

}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
