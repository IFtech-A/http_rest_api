package store

import "github.com/IFtech-A/http_rest_api/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
