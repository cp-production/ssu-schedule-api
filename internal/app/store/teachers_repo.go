package store

import (
	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type TeachersRepo struct {
	store *Store
}

func (r *TeachersRepo) Insert(t *model.Teachers) error {
	query := "INSERT INTO teachers VALUES (DEFAULT, $1)"
	if _, err := r.store.DB().Exec(query, t.FullName); err != nil {
		return err
	}

	return nil
}

func (r *TeachersRepo) SelectAll() (*[]model.Teachers, error) {
	rows, err := r.store.DB().Query("SELECT * FROM teachers")
	if err != nil {
		return nil, err
	}
	var teachersArray []model.Teachers
	for rows.Next() {
		t := &model.Teachers{}
		if err := rows.Scan(&t.FullName); err != nil {
			return nil, err
		}
		teachersArray = append(teachersArray, *t)
	}

	return &teachersArray, nil
}

func (r *TeachersRepo) Delete() error {
	query := "TRUNCATE TABLE teachers RESTART IDENTITY"
	if _, err := r.store.DB().Exec(query); err != nil {
		return err
	}
	return nil
}
