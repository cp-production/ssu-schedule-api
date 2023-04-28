package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Table interface {
	Insert(data interface{}) error
	Select(query string) ([]interface{}, error)
}

type Store struct {
	config               *Config
	db                   *sql.DB
	departmentsRepo      *DepartmentsRepo
	groupsRepo           *GroupsRepo
	studentsScheduleRepo *StudentsScheduleRepo
	subgroupsRepo        *SubgroupsRepo
	teachersRepo         *TeachersRepo
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() error {
	s.db.Close()
	return nil
}

func (s *Store) Departments() *DepartmentsRepo {
	if s.departmentsRepo != nil {
		return s.departmentsRepo
	}

	s.departmentsRepo = &DepartmentsRepo{
		store: s,	}
	return s.departmentsRepo
}

func (s *Store) Groups() *GroupsRepo {
	if s.groupsRepo != nil {
		return s.groupsRepo
	}

	s.groupsRepo = &GroupsRepo{
		store: s,
	}
	return s.groupsRepo
}

func (s *Store) StudentsSchedule() *StudentsScheduleRepo {
	if s.studentsScheduleRepo != nil {
		return s.studentsScheduleRepo
	}

	s.studentsScheduleRepo = &StudentsScheduleRepo{
		store: s,
	}
	return s.studentsScheduleRepo
}

func (s *Store) Subgroups() *SubgroupsRepo {
	if s.subgroupsRepo != nil {
		return s.subgroupsRepo
	}

	s.subgroupsRepo = &SubgroupsRepo{
		store: s,
	}
	return s.subgroupsRepo
}

func (s *Store) Teachers() *TeachersRepo {
	if s.teachersRepo != nil {
		return s.teachersRepo
	}

	s.teachersRepo = &TeachersRepo{
		store: s,
	}
	return s.teachersRepo
}
