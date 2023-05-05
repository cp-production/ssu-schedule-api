package parser

import (
	"strings"
	"fmt"
	"time"
	// "sync"

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
	if err := s.Subgroups().Delete(); err != nil {
		return err
	}
	return nil
}

// TODO: Test and fix it more (im sorry, Tracto)
// func ParseAll(s *store.Store) error {
// 	if err := ClearDB(s); err != nil {
// 		return err
// 	}
// 	departments, err := ParseDepartments(s)
// 	if err != nil {
// 		return err
// 	}

// 	var wg sync.WaitGroup
// 	for _, d := range *departments {
// 		wg.Add(1)
// 		go func(d model.Departments) error {
// 			defer wg.Done()

// 			start := time.Now()
// 			groups, err := ParseGroups(s, d.Url, d.Id)
// 			if err != nil {
// 				return err
// 			}
// 			for _, g := range *groups {
// 				if err := ParseStudentsSchedule(s, g.Id, g.EdForm, g.GroupNum, d.Url); err != nil {
// 					return err
// 				}
// 			}
// 			fmt.Printf("Parsed department %s in %v\n", d.Url, time.Since(start))
// 			return nil
// 		}(d)
// 	}

// 	wg.Wait()

// 	return nil
// }

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
			if err := ParseStudentsSchedule(s, g.Id, g.EdForm, g.GroupNum, d.Url); err != nil {
				return err
			}
		}
		fmt.Printf("Parsed department %s in %v\n", d.Url, time.Since(start))
	}
	if err := ParseSubgroups(s); err != nil {
		return err
	}


	return nil
}

func ParseDepartments(s *store.Store) (*[]model.Departments, error) {
	depList, err := tracto.GetDepartmentList()
	if err != nil {
		return nil, err
	}

	departmentsRepo := s.Departments()
	for i, rows := range depList.DepartmentsList {
		department := model.Departments{
			Id:        i,
			ShortName: rows.ShortName,
			FullName:  rows.FullName,
			Url:       rows.Url,
		}
		if err := departmentsRepo.Insert(&department); err != nil {
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
		group := model.Groups{
			EdForm:   rows[0],
			GroupNum: rows[1],
			DepId:    depID,
		}
		if err := groupsRepo.Insert(&group); err != nil {
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

func ParseSubgroups(s *store.Store) error {
	schedule, err := s.StudentsSchedule().SelectAll()
	if err != nil {
		return err
	}
	subgroupSet := make(map[model.Subgroups]bool)
	for _, lesson := range *schedule {
		subgroup := model.Subgroups{
			SubgroupName: lesson.SubgroupName,
			GroupId:      lesson.GroupId,
		}
		if _, exist := subgroupSet[subgroup]; !exist {
			subgroupSet[subgroup] = true
		}
	}

	subgroupsRepo := s.Subgroups()
	for subgroup := range subgroupSet {
		err := subgroupsRepo.Insert(&subgroup)
		if err != nil {
			return err
		}
	}

	return nil
}
