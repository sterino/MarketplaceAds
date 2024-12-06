package influencer

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrorNotFound            = errors.New("error not found")
	ErrorEmailConflict       = errors.New("email already exists")
	ErrorInvalidName         = errors.New("name is invalid or empty")
	ErrorInvalidEmail        = errors.New("email format is invalid")
	ErrorInvalidPlatforms    = errors.New("platforms must contain at least one valid URL")
	ErrorInvalidCategory     = errors.New("category is required")
	ErrorInvalidFollowers    = errors.New("followers must be greater than zero")
	ErrorInvalidPricePerPost = errors.New("price per post must be greater than zero")
)

type RegisterRequest struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	PhoneNumber    string   `json:"phone_number"`
	Platforms      []string `json:"platforms"`
	FollowersCount int      `json:"followers_count"`
	Category       string   `json:"category"`
	Bio            string   `json:"bio"`
	Address        string   `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerified   bool      `json:"email_verified"`
	PhoneNumber     string    `json:"phone_number"`
	AccountVerified bool      `json:"account_verified"`
	AccountType     string    `json:"account_type"`
	Platforms       []string  `json:"platforms"`
	FollowersCount  int       `json:"followers_count"`
	Category        string    `json:"category"`
	Bio             string    `json:"bio"`
	Address         string    `json:"address,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
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
		ID:             entity.ID,
		Name:           entity.Name,
		Email:          entity.Email,
		PhoneNumber:    entity.PhoneNumber,
		Platforms:      entity.Platforms,
		FollowersCount: entity.FollowersCount,
		Category:       entity.Category,
		Bio:            entity.Bio,
		Address:        entity.Address,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}

func (r *RegisterRequest) Validate() error {
	if r.Name == "" {
		return ErrorInvalidName
	}
	if !isValidEmail(r.Email) {
		return ErrorInvalidEmail
	}
	if len(r.Platforms) == 0 || !areValidURLs(r.Platforms) {
		return ErrorInvalidPlatforms
	}
	if r.Category == "" {
		return ErrorInvalidCategory
	}
	if r.FollowersCount <= 0 {
		return ErrorInvalidFollowers
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func areValidURLs(urls []string) bool {
	re := regexp.MustCompile(`^https?://[\w\-]+(\.[\w\-]+)+[/#?]?.*$`)
	for _, url := range urls {
		if !re.MatchString(url) {
			return false
		}
	}
	return true
}
