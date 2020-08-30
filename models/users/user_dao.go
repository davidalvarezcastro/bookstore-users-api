package users

import (
	"fmt"

	"github.com/davidalvarezcastro/bookstore-users-api/datasources/mysql/usersdb"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
	mysqlutils "github.com/davidalvarezcastro/bookstore-users-api/utils/mysql"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status = ?"
)

var (
	userDB = make(map[int64]*User)
)

// Get returns the user by given userId
func (u *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

// Save saves a user in the database
func (u *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.DateCreated)
	if err != nil {
		fmt.Println(err)
		return mysqlutils.ParseError(err)
	}
	// insertResult, err := stmt.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	u.ID = userID
	return nil
}

// Update updates user info in the database
func (u *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

// Delete removes an user from the database
func (u *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.ID); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

// FindByStatus searches an user by status in the database
func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); err != nil {
			return nil, mysqlutils.ParseError(err)
		}

		results = append(results, u)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
