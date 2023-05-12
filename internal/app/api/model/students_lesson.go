package model

type StudentsLesson struct {
	// TODO: id -> group_id
	GroupId      int    `json:"id"`
	DayNum       int    `json:"day_num"`
	WeekType     string `json:"week_type"`
	LessonType   string `json:"lesson_type"`
	LessonName   string `json:"lesson_name"`
	LessonNumber int    `json:"lesson_number"`
	Teacher      string `json:"teacher"`
	LessonPlace  string `json:"lesson_place"`
	SubgroupName string `json:"subgroup_name"`
}
