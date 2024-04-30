package users

import (
	"fmt"
	"time"

	"rainmore.com.au/rest-api/datasources/postgresql/users_db"
	"rainmore.com.au/rest-api/domain/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func Exist(id int64) (bool, *errors.RestError) {
	users_db.Ping()

	stmt, err := users_db.DBClient.Prepare("SELECT COUNT(id) FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return false, errors.NewError(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	exists := false
	if err := row.Scan(&exists); err != nil {
		fmt.Println(err)
		return false, errors.NewError(err)
	}

	return exists, nil
}

func Get(user *User) *errors.RestError {
	users_db.Ping()

	stmt, err := users_db.DBClient.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	row := stmt.QueryRow(user.Id)

	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.CreatedAt); err != nil {
		fmt.Println(err)
		return errors.NewError(err)
	}

	return nil
}

func Save(user *User) *errors.RestError {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	var dateOfBirthStr string

	if user.DateOfBirth != nil {
		dateOfBirthStr = user.DateOfBirth.ToString()
	}

	stmt, err := users_db.DBClient.Prepare(
		`INSERT INTO users (first_name, last_name, email, date_of_birth, created_at)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id`)

	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}

	row := stmt.QueryRow(
		user.FirstName,
		user.LastName,
		user.Email,
		dateOfBirthStr,
		user.CreatedAt)

	if err := row.Scan(&user.Id); err != nil {
		fmt.Println(err)
		return errors.NewError(err)
	}

	return Get(user)
}

func Update(user *User) *errors.RestError {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	var dateOfBirthStr string

	if user.DateOfBirth != nil {
		dateOfBirthStr = user.DateOfBirth.ToString()
	}

	stmt, err := users_db.DBClient.Prepare(
		`UPDATE users SET
		first_name = $1, 
		last_name = $2, 
		email = $3, 
		date_of_birth = $4
		WHERE id = $5
		RETURNING id`)

	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError(err.Error())
	}

	row := stmt.QueryRow(
		user.FirstName,
		user.LastName,
		user.Email,
		dateOfBirthStr,
		user.Id)

	if err := row.Scan(&user.Id); err != nil {
		fmt.Println(err)
		return errors.NewError(err)
	}

	return Get(user)
}

func Delete(id int64) (bool, *errors.RestError) {
	users_db.Ping()

	stmt, err := users_db.DBClient.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return false, errors.NewError(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return false, errors.NewError(err)
	}

	c, err := res.RowsAffected()

	if err != nil {
		return false, errors.NewError(err)
	}

	return c == 1, nil
}
