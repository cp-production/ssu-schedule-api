package parser

import (
	"fmt"
	"strings"

	"time"

	"github.com/cp-production/ssu-schedule-api/internal/app/api/model"
	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/cp-production/ssu-schedule-api/internal/app/tracto"
)

func ClearDB(s *store.Store) error {
	if err := s.StudentsSchedule().Delete(); err != nil {
		return err
	}
	if err := s.Groups().Delete(); err != nil {
		return err
	}
	if err := s.Departments().Delete(); err != nil {
		return err
	}
	return nil
}

func ParseAll(s *store.Store) error {
	if err := ClearDB(s); err != nil {
		return err
	}

	departments, err := ParseDepartments(s)
	if err != nil {
		return err
	}

	for _, d := range *departments {
		start := time.Now()
		groups, err := ParseGroups(s, d.Url, d.Id)
		if err != nil {
			return err
		}
		for _, g := range *groups {
			err := ParseStudentsSchedule(s, g.Id, g.EdForm, g.GroupNum, d.Url)
			if err != nil {
				return err
			}
		}
		end := time.Now()
		fmt.Println(d.Url, end.Sub(start))
	}

	return nil
}

func ParseDepartments(s *store.Store) (*[]model.Departments, error) {
	depList, err := tracto.GetDepartmentList()
	if err != nil {
		return nil, err
	}

	departmentsRepo := s.Departments()
	for i, rows := range depList {
		department := model.Departments{Id: i, Url: rows[0], FullName: rows[1]}
		err := departmentsRepo.Insert(&department)
		if err != nil {
			return nil, err
		}
	}

	processedDepartments, err := departmentsRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return processedDepartments, nil
}

func ParseGroups(s *store.Store, url string, depID int) (*[]model.Groups, error) {

	groupList, err := tracto.GetGroupList(url)
	if err != nil {
		return nil, err
	}

	groupsRepo := s.Groups()
	for _, rows := range groupList {
		group := model.Groups{EdForm: rows[0], GroupNum: rows[1], DepId: depID}
		err := groupsRepo.Insert(&group)
		if err != nil {
			return nil, err
		}
	}

	processedGroups, err := groupsRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return processedGroups, nil
}

func ParseStudentsSchedule(s *store.Store, groupId int, educationForm string, groupNum string, department string) error {
	var tractoEdForm string
	if educationForm == "do" {
		tractoEdForm = "full"
	} else {
		tractoEdForm = "extramural"
	}
	schedule := tracto.GetSchedule(tractoEdForm, department, groupNum)

	scheduleRepo := s.StudentsSchedule()

	for _, rows := range schedule.Lessons {
		teacherFullName := strings.TrimSpace(rows.Teacher.Surname) + " " + strings.TrimSpace(rows.Teacher.Name) + " " + strings.TrimSpace(rows.Teacher.Patronymic)
		lesson := model.StudentsLesson{
			GroupId:      groupId,
			DayNum:       rows.Day.DayNumber,
			WeekType:     rows.WeekType,
			LessonType:   rows.LessonType,
			LessonName:   rows.Name,
			Teacher:      teacherFullName,
			LessonPlace:  rows.Place,
			SubgroupName: rows.SubGroup}
		err := scheduleRepo.Insert(&lesson)
		if err != nil {
			return err
		}
	}

	return nil
}
