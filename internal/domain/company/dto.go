package company

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrorNotFound        = errors.New("error not found")
	ErrorEmailConflict   = errors.New("email already exists")
	ErrorInvalidName     = errors.New("name is invalid or empty")
	ErrorInvalidEmail    = errors.New("email format is invalid")
	ErrorInvalidPassword = errors.New("password must be at least 8 characters")
)

type RegisterRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, entity := range data {
		res = append(res, ParseFromEntity(entity))
	}
	return
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:          entity.ID,
		Name:        entity.Name,
		Email:       entity.Email,
		PhoneNumber: entity.PhoneNumber,
		Address:     entity.Address,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (r *RegisterRequest) Validate() error {
	if r.Name == "" {
		return ErrorInvalidName
	}
	if !isValidEmail(r.Email) {
		return ErrorInvalidEmail
	}
	if len(r.Password) < 8 {
		return ErrorInvalidPassword
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
