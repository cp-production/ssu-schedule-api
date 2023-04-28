package store

import "github.com/cp-production/ssu-schedule-api/internal/app/api/model"

type SubgroupsRepo struct {
	store *Store
}

func (r *SubgroupsRepo) Insert(s *model.Subgroups) error {
	query := "INSERT INTO subgroups VALUES($1, $2)"
	if _, err := r.store.db.Exec(query, s.GroupId, s.SubgroupName); err != err {
		return err
	}
	return nil
}

func (r *SubgroupsRepo) SelectByGroup(query_id string) (*model.Subgroups, error) {
	query := "SELECT * FROM subgroups WHERE group_id = $1"
	row := r.store.db.QueryRow(query, query_id)

	sub := &model.Subgroups{}
	if err := row.Scan(&sub.SubgroupName); err != nil {
		return nil, err
	}
	return sub, nil
}

func (r *SubgroupsRepo) Delete() error {
	query := "TRUNCATE TABLE subgroups RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
