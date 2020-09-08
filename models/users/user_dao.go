package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/davidalvarezcastro/bookstore-users-api/datasources/mysql/usersdb"
	"github.com/davidalvarezcastro/bookstore-users-api/logger"
	mysqlutils "github.com/davidalvarezcastro/bookstore-users-api/utils/mysql"
	errorsutils "github.com/davidalvarezcastro/bookstore-utils-go/rest_errors"
)

const (
	queryInsertUser                 = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                    = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id = ?;"
	queryUpdateUser                 = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser                 = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus           = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status = ?"
	queryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE email = ? AND password = ? AND status = ?"
)

var (
	userDB = make(map[int64]*User)
)

// Get returns the user by given userId
func (u *User) Get() *errorsutils.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
		// return mysqlutils.ParseError(err)
	}

	return nil
}

// Save saves a user in the database
func (u *User) Save() *errorsutils.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.DateCreated)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	// insertResult, err := stmt.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creatin a new user", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}

	u.ID = userID
	return nil
}

// Update updates user info in the database
func (u *User) Update() *errorsutils.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID); err != nil {
		logger.Error("error when trying to update user", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}

	return nil
}

// Delete removes an user from the database
func (u *User) Delete() *errorsutils.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}

	return nil
}

// FindByStatus searches an user by status in the database
func (u *User) FindByStatus(status string) ([]User, *errorsutils.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); err != nil {
			logger.Error("error when trying to scan user row into user struct ", err)
			return nil, errorsutils.NewInternalServerError("database error", errors.New("database error"))
		}

		results = append(results, u)
	}

	if len(results) == 0 {
		return nil, errorsutils.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

// FindByEmailAndPassword returns the user by given email and password
func (u *User) FindByEmailAndPassword() *errorsutils.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Email, u.Password, StatusActive)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysqlutils.ErrorNoRows) {
			return errorsutils.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", err)
		return errorsutils.NewInternalServerError("database error", errors.New("database error"))
		// return mysqlutils.ParseError(err)
	}

	return nil
}
