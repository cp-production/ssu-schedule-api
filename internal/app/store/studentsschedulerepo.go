package store

import (
	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type StudentsScheduleRepo struct {
	store *Store
}

func (r *StudentsScheduleRepo) Insert(s *model.StudentsLesson) error {

	query := "INSERT INTO studentsSchedule VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)"
	if _, err := r.store.db.Exec(query, s.GroupId, s.DayNum, s.WeekType, s.LessonType,
		s.LessonName, s.Teacher, s.LessonPlace, s.SubgroupName); err != nil {
		return err
	}
	return nil
}

func (r *StudentsScheduleRepo) Select() (*model.StudentsLesson, error) {
	return nil, nil
}

func (r *StudentsScheduleRepo) Delete() error {
	query := "TRUNCATE TABLE studentsSchedule RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
