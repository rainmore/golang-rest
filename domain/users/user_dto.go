package users

import (
	"strings"
	"time"

	"rainmore.com.au/rest-api/domain/date_time"
	"rainmore.com.au/rest-api/domain/errors"
)

type User struct {
	Id          int64               `json:"id"`
	FirstName   string              `json:"first_name"`
	LastName    string              `json:"last_name"`
	Email       string              `json:"email"`
	DateOfBirth *date_time.DateOnly `json:"date_of_birth"`
	CreatedAt   time.Time           `json:"created_at"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

func (user *User) CopyTo(newUser *User) {
	newUser.Id = user.Id
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Email = user.Email
	newUser.DateOfBirth = user.DateOfBirth
	newUser.CreatedAt = user.CreatedAt
}
