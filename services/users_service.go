package services

import (
	"fmt"
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

func UpdateUser(user users.User) *errors.RestError {
	if err := user.Validate(); err != nil {
		return err
	}

	exists, err := users.Exist(user.Id)

	if err != nil || !exists {
		return errors.NewNotFundError(fmt.Sprintf("User with %d does not exist", user.Id))
	}

	if err := users.Update(&user); err != nil {
		return err
	}

	return nil
}

func DeleteUser(user users.User) *errors.RestError {
	if err := user.Validate(); err != nil {
		return err
	}

	exists, err := users.Exist(user.Id)

	if err != nil || !exists {
		return errors.NewNotFundError(fmt.Sprintf("User with %d does not exist", user.Id))
	}

	res, err := users.Delete(user.Id)

	if err != nil {
		return err
	}

	if !res {
		return errors.NewBadRequestError("No row was deleted")
	}

	return nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId, CreatedAt: time.Now()}

	if err := users.Get(result); err != nil {
		return nil, err
	}
	return result, nil
}

func Exist(userId int64) (bool, *errors.RestError) {
	return users.Exist(userId)
}
