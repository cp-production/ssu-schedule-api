package model

type TeachersSchedule struct {
	Id           int `json:"id"`
	DayNum       string `json:"day_num"`
	WeekType     string `json:"week_type"`
	LessonType   string `json:"lesson_type"`
	LessonName   string `json:"lesson_name"`
	GroupNum     string `json:"group_num"`
	LessonPlace  string `json:"lesson_place"`
	SubgroupName string `json:"subgroup_name"`
}
