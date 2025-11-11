package usecase

import (
	"errors"
	"hexa-fiber-gorm/internal/modules/user/domain"
	ports "hexa-fiber-gorm/internal/modules/user/port"
)

type userUsecase struct {
	repo ports.UserRepository
}

func NewUserUsecase(repo ports.UserRepository) ports.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &domain.User{Name: name, Email: email}
	if err := u.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserByID(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) GetAllUsers() ([]domain.User, error) {
	return u.repo.FindAll()
}
