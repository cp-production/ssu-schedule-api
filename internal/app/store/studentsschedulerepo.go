package store

import "github.com/cp-production/ssu-schedule-api/internal/app/api/model"

type StudentsScheduleRepo struct {
	store *Store
}

func (r *StudentsScheduleRepo) Insert(s *model.StudentsSchedule) error {

	query := "INSERT INTO studentsSchedule VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	if _, err := r.store.db.Exec(query, s.DayNum, s.GroupId, s.LessonName,
		s.LessonPlace, s.LessonType, s.SubgroupName,
		s.Teacher, s.WeekType); err != nil {
		return err
	}
	return nil
}

func (r *StudentsScheduleRepo) Select() (*model.StudentsSchedule, error) {
	return nil, nil
}

func (r *StudentsScheduleRepo) Delete() error {
	query := "TRUNCATE TABLE studentsSchedule RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
