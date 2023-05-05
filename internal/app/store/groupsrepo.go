package store

import (
	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type GroupsRepo struct {
	store *Store
}

func (r *GroupsRepo) Insert(g *model.Groups) error {
	query := "INSERT INTO groups VALUES (DEFAULT, $1, $2, $3)"
	if _, err := r.store.db.Exec(query, g.EdForm, g.GroupNum, g.DepId); err != nil {
		return err
	}

	return nil
}

func (r *GroupsRepo) SelectByDepartments(ed_form string, dep_url string) (*[]model.Groups, error) {
	query := "SELECT * FROM groups WHERE edForm = $1 AND dep_id = (SELECT id FROM departments WHERE url = $2)"
	rows, err := r.store.db.Query(query, ed_form, dep_url)

	if err != nil {
		return nil, err
	}

	var groups []model.Groups
	for rows.Next() {
		group := &model.Groups{}
		if err := rows.Scan(&group.Id, &group.EdForm, &group.GroupNum, &group.DepId); err != nil {
			return nil, err
		}
		groups = append(groups, *group)
	}
	return &groups, nil
}

func (r *GroupsRepo) SelectId(edForm string, groupNum string, url string) (int, error) {
	query := "SELECT id FROM groups WHERE edForm = $1 AND groupNum = $2 AND dep_id = (SELECT id FROM departments WHERE url = $3)"
	row := r.store.db.QueryRow(query, edForm, groupNum, url)

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *GroupsRepo) SelectAll() (*[]model.Groups, error) {
	rows, err := r.store.db.Query("SELECT * FROM groups")
	if err != nil {
		return nil, err
	}
	var groupsArray []model.Groups
	for rows.Next() {
		group := &model.Groups{}
		if err := rows.Scan(&group.Id, &group.EdForm, &group.GroupNum, &group.DepId); err != nil {
			return nil, err
		}
		groupsArray = append(groupsArray, *group)
	}

	return &groupsArray, nil
}

func (r *GroupsRepo) Delete() error {
	query := "TRUNCATE TABLE groups RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
