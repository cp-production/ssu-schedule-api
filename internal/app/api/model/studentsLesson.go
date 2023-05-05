package model

type StudentsLesson struct {
	GroupId      int `json:"id"`
	DayNum       int `json:"dayNum"`
	WeekType     string `json:"weekType"`
	LessonType   string `json:"lessonType"`
	LessonName   string `json:"lessonName"`
	Teacher      string `json:"teacher"`
	LessonPlace  string `json:"lessonPlace"`
	SubgroupName string `json:"subgroupName"`
}
