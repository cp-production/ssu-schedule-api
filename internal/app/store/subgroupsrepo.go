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

func (r *SubgroupsRepo) Select() (*model.Subgroups, error) {
	return nil, nil
}

func (r *SubgroupsRepo) Delete() error {
	query := "TRUNCATE TABLE subgroups"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
