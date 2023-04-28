package store

import (
	"fmt"

	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type DepartmentsRepo struct {
	store *Store
}

func (r *DepartmentsRepo) Insert(d *model.Departments) error {

	query := "INSERT INTO departments VALUES (DEFAULT, $1, $2)"
	if _, err := r.store.db.Exec(query, d.Url, d.FullName); err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func (r *DepartmentsRepo) SelectAll() (*[]model.Departments, error) {
	rows, err := r.store.db.Query("SELECT * FROM departments")
	if err != nil {
		return nil, err
	}
	var departmentsArray []model.Departments
	for rows.Next() {
		dep := &model.Departments{}
		if err := rows.Scan(&dep.Url, &dep.FullName); err != nil {
			return nil, err
		}
		departmentsArray = append(departmentsArray, *dep)
	}

	return &departmentsArray, nil
}

func (r *DepartmentsRepo) SelectById(query_id string) (*model.Departments, error) {
	query := "SELECT * FROM departments WHERE id = $1"
	row := r.store.db.QueryRow(query, query_id)

	dep := &model.Departments{}
	if err := row.Scan(&dep.Url, &dep.FullName); err != nil {
		return nil, err
	}
	return dep, nil

}
func (r *DepartmentsRepo) Delete() error {
	query := "TRUNCATE TABLE departments RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}