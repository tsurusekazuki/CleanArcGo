package usecase

import "github.com/tsurusekazuki/cleanArcGo/src/app/domain"

type UserRepository interface {
	Store(user domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}