package store

import (
	"fmt"

	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
)

type StudentsScheduleRepo struct {
	store *Store
}

func (r *StudentsScheduleRepo) Insert(s *model.StudentsLesson) error {
	query := "INSERT INTO students_schedule VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9)"
	if _, err := r.store.DB().Exec(query, s.GroupId, s.DayNum, s.WeekType, s.LessonType,
		s.LessonName, s.LessonNumber, s.Teacher, s.LessonPlace, s.SubgroupName); err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func (r *StudentsScheduleRepo) Select(educationForm string, departmentUrl string, groupNum string) (*[]model.StudentsLesson, error) {
	query := "SELECT * FROM students_schedule WHERE group_id = (SELECT id FROM groups WHERE education_form = $1 AND department_id = (SELECT id FROM departments WHERE url = $2) AND group_num = $3)"
	rows, err := r.store.DB().Query(query, educationForm, departmentUrl, groupNum)
	if err != nil {
		return nil, err
	}
	var lessons []model.StudentsLesson
	for rows.Next() {
		lesson := &model.StudentsLesson{}
		// TODO: Fix this
		var id int
		if err := rows.Scan(&id, &lesson.GroupId, &lesson.DayNum, &lesson.WeekType, &lesson.LessonType, &lesson.LessonName, &lesson.LessonNumber, &lesson.Teacher, &lesson.LessonPlace, &lesson.SubgroupName); err != nil {
			return nil, err
		}
		lessons = append(lessons, *lesson)
	}
	return &lessons, nil
}

func (r *StudentsScheduleRepo) SelectAll() (*[]model.StudentsLesson, error) {
	query := "SELECT * FROM students_schedule"
	rows, err := r.store.DB().Query(query)
	if err != nil {
		return nil, err
	}
	var lessons []model.StudentsLesson
	for rows.Next() {
		lesson := &model.StudentsLesson{}
		// TODO: Change this
		var id int
		if err := rows.Scan(&id, &lesson.GroupId, &lesson.DayNum, &lesson.WeekType, &lesson.LessonType, &lesson.LessonName, &lesson.LessonNumber, &lesson.Teacher, &lesson.LessonPlace, &lesson.SubgroupName); err != nil {
			return nil, err
		}
		lessons = append(lessons, *lesson)
	}
	return &lessons, nil
}

func (r *StudentsScheduleRepo) Delete() error {
	query := "TRUNCATE TABLE students_schedule RESTART IDENTITY"
	if _, err := r.store.DB().Exec(query); err != nil {
		return err
	}
	return nil
}
