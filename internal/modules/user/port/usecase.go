package ports

import "hexa-fiber-gorm/internal/modules/user/domain"

type UserUsecase interface {
	CreateUser(name, email string) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
}
