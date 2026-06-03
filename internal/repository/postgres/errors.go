package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrDuplicateEmail         = errors.New("duplicate email")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserCredentialNotFound = errors.New("user credential not found")
)

const (
	PgUniqueViolation     = "23505"
	PgForeignKeyViolation = "23503"
	PgCheckViolation      = "23514"
)

func mapPostgresError(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code {

	case PgUniqueViolation:
		switch pgErr.ConstraintName {
		case ConstraintUsersEmailUnique:
			return ErrDuplicateEmail
		}
	}

	return err
}
