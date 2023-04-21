package store

import "github.com/cp-production/ssu-schedule-api/internal/app/api/model"

type GroupsRepo struct {
	store *Store
}

func (r *GroupsRepo) Insert(g *model.Groups) error {
	query := "INSERT INTO departments VALUES ($1, $2, $3)"
	if _, err := r.store.db.Exec(query, g.EdForm, g.GroupNum, g.DepId); err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) Select() (*model.Groups, error) {

	return nil, nil
}

func (r *GroupsRepo) Delete() error {
	query := "TRUNCATE TABLE groups"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
