package services

import (
	"time"

	"rainmore.com.au/rest-api/domain/errors"
	"rainmore.com.au/rest-api/domain/users"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := users.Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId, CreatedAt: time.Now()}

	if err := users.Get(result); err != nil {
		return nil, err
	}
	return result, nil
}
