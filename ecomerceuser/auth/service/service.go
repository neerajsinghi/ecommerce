package service

import (
	"errors"
	"time"

	"ecommerceuser/model"
	"ecommerceuser/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID       int
	Username string
	Password string
}

// AuthService represents the authentication service
type AuthService struct {
	repo        repository.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

// NewAuthService creates a new AuthService
func NewAuthService(repo repository.UserRepository, jwtSecret string, tokenExpiry time.Duration) *AuthService {
	return &AuthService{
		repo:        repo,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

// Register registers a new user
func (s *AuthService) Register(user *model.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	//check user
	_, err = s.repo.Login(user.Email, user.Password)
	if err == nil {
		return "", errors.New("user already exists")
	}
	pass := user.Password
	user.Password = string(hashedPassword)
	err = s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	return s.Login(user.Email, pass)
}

// Login logs in a user and returns a JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.Login(email, password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(s.tokenExpiry).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logout logs out a user (this is a placeholder as JWT tokens are stateless)
func (s *AuthService) Logout(token string) error {
	// Invalidate the token if using a token store, otherwise this is a no-op
	return nil
}
