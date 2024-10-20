package repository

import (
	"ecommerceuser/model"
	"errors"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255;unique"`
}

// UserRepository provides access to the user storage
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	user := &model.User{}
	if err := r.db.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *model.User) error {
	tx := r.db.Create(user)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) Login(username, password string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("email = ?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *model.User) error {
	if err := r.db.Model(User{}).Where("id=?", user.ID).Updates(User{Name: user.Name}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database
func (r *UserRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
