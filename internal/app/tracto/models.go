package tracto

const TractoUri = "https://scribaproject.space/api/v1.0"
const SsuUri = "https://www.sgu.ru/schedule"

type lessonTime struct {
	Id           int `json:"id"`
	LessonNumber int `json:"lessonNumber"`
	HourStart    int `json:"hourStart"`
	MinuteStart  int `json:"minuteStart"`
	HourEnd      int `json:"hourEnd"`
	MinuteEnd    int `json:"minuteEnd"`
}

type teacher struct {
	Id         int    `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
}

type Lessons []struct {
	Id               int          `json:"id"`
	Name             string       `json:"name"`
	Place            string       `json:"place"`
	Department       Department   `json:"department"`
	StudentGroup     studentGroup `json:"studentGroup"`
	SubGroup         string       `json:"subGroup"`
	Day              Day          `json:"day"`
	LessonTime       lessonTime   `json:"lessonTime"`
	Teacher          teacher      `json:"teacher"`
	WeekType         string       `json:"weekType"`
	LessonType       string       `json:"lessonType"`
	UpdatedTimestamp int          `json:"updatedTimestamp"`
	BeginTimestamp   int          `json:"beginTimestamp"`
	EndTimestamp     int          `json:"endTimestamp"`
}

type Department struct {
	Id        int    `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Url       string `json:"url"`
}

type Departments struct {
	DepartmentsList []Department
}
type studentGroup struct {
	Id             int        `json:"id"`
	GroupNumber    string     `json:"groupNumber"`
	GroupNumberRus string     `json:"groupNumberRus"`
	Department     Department `json:"department"`
	EducationForm  string     `json:"educationForm"`
	GroupType      string     `json:"groupType"`
}

type Day struct {
	Id        int    `json:"id"`
	DayNumber int    `json:"dayNumber"`
	WeekDay   string `json:"weekDay"`
}

type weekShift struct {
	Id         int    `json:"id"`
	Shift      int    `json:"shift"`
	Department string `json:"department"`
}

type Schedule struct {
	Lessons      Lessons      `json:"lessons"`
	StudentGroup studentGroup `json:"studentGroup"`
	Day          Day          `json:"day"`
	WeekShift    weekShift    `json:"weekShift"`
}
