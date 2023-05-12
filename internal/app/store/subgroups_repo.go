package store

import (
	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type SubgroupsRepo struct {
	store *Store
}

func (r *SubgroupsRepo) Insert(s *model.Subgroups) error {
	query := "INSERT INTO subgroups VALUES($1, $2)"
	if _, err := r.store.DB().Exec(query, s.GroupId, s.SubgroupName); err != err {
		return err
	}
	return nil
}

func (r *SubgroupsRepo) SelectByGroup(education_form string, department_url string, group_num string) (*[]model.Subgroups, error) {
	query := "SELECT * FROM subgroups WHERE group_id = (SELECT id FROM groups WHERE education_form = $1 AND group_num = $2 AND department_id = (SELECT id FROM departments WHERE url = $3))"
	rows, err := r.store.DB().Query(query, education_form, group_num, department_url)
	if err != nil {
		return nil, err
	}

	var subgroupsArray []model.Subgroups
	for rows.Next() {
		subgroup := &model.Subgroups{}
		// TODO: Change this
		var i int
		if err := rows.Scan(&i, &subgroup.SubgroupName, &subgroup.GroupId); err != nil {
			return nil, err
		}
		subgroupsArray = append(subgroupsArray, *subgroup)
	}

	return &subgroupsArray, nil
}

func (r *SubgroupsRepo) Delete() error {
	query := "TRUNCATE TABLE subgroups RESTART IDENTITY"
	if _, err := r.store.DB().Exec(query); err != nil {
		return err
	}
	return nil
}
