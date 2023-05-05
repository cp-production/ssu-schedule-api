package model

type TeachersSchedule struct {
	Id           int `json:"id"`
	DayNum       string `json:"dayNum"`
	WeekType     string `json:"weekType"`
	LessonType   string `json:"lessonType"`
	LessonName   string `json:"lessonName"`
	GroupNum     string `json:"groupNum"`
	LessonPlace  string `json:"lessonPlace"`
	SubgroupName string `json:"subgroupName"`
}
