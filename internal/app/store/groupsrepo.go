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

func (r *GroupsRepo) SelectByDepartments(query_id string) (*model.Groups, error) {
	query := "SELECT * FROM groups WHERE dep_id = $1"
	row := r.store.db.QueryRow(query, query_id)

	group := &model.Groups{}
	if err := row.Scan(&group.EdForm, &group.GroupNum); err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupsRepo) Delete() error {
	query := "TRUNCATE TABLE groups RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
