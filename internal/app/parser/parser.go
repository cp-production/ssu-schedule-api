package parser

import (
	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/cp-production/ssu-schedule-api/internal/app/tracto"
)

func ParseDepartments(s *store.Store) error {
	depList, err := tracto.GetDepartmentList()

	if err != nil {
		return err
	} else {
		departmentsRepo := s.Departments()
		err := departmentsRepo.Delete()
		if err != nil {
			return nil
		}

		for _, rows := range depList {
			department := model.Departments{Url: rows[0], FullName: rows[1]}
			departmentsRepo.Insert(&department)
		}
	}

	return nil
}
