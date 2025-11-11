package ports

import "hexa-fiber-gorm/internal/modules/user/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]domain.User, error)
}
