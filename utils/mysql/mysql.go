package mysqlutils

import (
	"errors"
	"strings"

	errorsutils "github.com/davidalvarezcastro/bookstore-utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	// ErrorNoRows is the message returned by the database
	ErrorNoRows = "no rows in result set"
)

// ParseError checks and return differents types of RestErr depends on the error receives
func ParseError(err error) *errorsutils.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errorsutils.NewNotFoundError("no record matching given id")
		}
		return errorsutils.NewInternalServerError("error parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return errorsutils.NewBadRequestError("invalid data")
	}

	return errorsutils.NewInternalServerError("error processing request", errors.New("database error"))
}
