package users

import (
	"fmt"
	"time"

	"rainmore.com.au/rest-api/domain/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func Get(user *User) *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFundError(fmt.Sprintf("user %d not found", user.Id))
	}

	result.CopyTo(user)
	return nil
}

func Save(user *User) *errors.RestError {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	usersDB[user.Id] = user
	return nil
}
