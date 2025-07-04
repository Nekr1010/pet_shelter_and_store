package repository

import (
	"database/sql"
	"errors"
	"pet_shelter_and_store/internal/errs"
)

func translateError(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return errs.ErrNotFound
	} else {
		return err
	}
}
