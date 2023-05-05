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

func (r *StudentsScheduleRepo) Select(educationForm string, departmentUrl string, groupNum string) (*[]model.StudentsLesson, error) {

	query := "SELECT * FROM studentsschedule WHERE group_id = (SELECT group_id FROM groups WHERE edForm = $1 AND dep_id = (SELECT id FROM departments WHERE url = $2) AND groupNum = $3)"
	rows, err := r.store.db.Query(query, educationForm, departmentUrl, groupNum)
	if err != nil {
		return nil, err
	}
	var lessons []model.StudentsLesson
	for rows.Next() {
		lesson := &model.StudentsLesson{}
		// TODO: Fix this
		var id int
		if err := rows.Scan(&id, &lesson.GroupId, &lesson.DayNum, &lesson.WeekType, &lesson.LessonType, &lesson.LessonName, &lesson.Teacher, &lesson.LessonPlace, &lesson.SubgroupName); err != nil {
			return nil, err
		}
		lessons = append(lessons, *lesson)
	}
	return &lessons, nil
}

func (r *StudentsScheduleRepo) SelectAll() (*[]model.StudentsLesson, error) {

	query := "SELECT * FROM studentsschedule"
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	var lessons []model.StudentsLesson
	for rows.Next() {
		lesson := &model.StudentsLesson{}
		// TODO: Change this
		var id int
		if err := rows.Scan(&id, &lesson.GroupId, &lesson.DayNum, &lesson.WeekType, &lesson.LessonType, &lesson.LessonName, &lesson.Teacher, &lesson.LessonPlace, &lesson.SubgroupName); err != nil {
			return nil, err
		}
		lessons = append(lessons, *lesson)
	}
	return &lessons, nil
}

func (r *StudentsScheduleRepo) Delete() error {
	query := "TRUNCATE TABLE studentsSchedule RESTART IDENTITY"
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
