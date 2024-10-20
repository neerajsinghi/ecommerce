package service

import (
	"context"
	"ecommerceuser/model"
	"ecommerceuser/repository"
	"errors"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// update user
func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	if user.ID == 0 {
		return errors.New("user ID is required")
	}
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

// delete user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("user ID is required")
	}

	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
