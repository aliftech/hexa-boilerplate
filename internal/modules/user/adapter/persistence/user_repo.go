package persistence

import (
	"hexa-fiber-gorm/internal/modules/user/domain"
	ports "hexa-fiber-gorm/internal/modules/user/port"
	"hexa-fiber-gorm/pkg/db"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDB *db.GORM) ports.UserRepository {
	return &userRepository{db: gormDB.DB()}
}

func (r *userRepository) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var u domain.User
	err := r.db.First(&u, id).Error
	return &u, err
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}
