package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary get a list of departments
// @ID get-departments-list
// @Tags departments
// @Description Retrieves SSU departments' list
// @Produce json
// @Success 200 {array} model.Departments
// @Router /departments [get]
func (s *server) handleDepartments() http.HandlerFunc {
	d, _ := s.store.Departments().SelectAll()

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, *d)
	}
}

// @Summary get a list of groups of a certain department
// @Tags groups
// @Description Retrieves groups' list based on department and education form
// @ID get-groups-list
// @Param education_form path string true "Education form, e.g. `do`"
// @Param department path string true "Department URL, e.g. `knt` for CSIT department"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Groups
// @Router /{education_form}/{department}/groups [get]
func (s *server) handleGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		educationForm := vars["ed_form"]
		departmentUrl := vars["dep_url"]
		d, err := s.store.Groups().SelectByDepartments(educationForm, departmentUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if d == nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}
		s.respond(w, http.StatusOK, *d)
	}
}

// @Summary get the schedule of students for a particular group
// @Tags schedule
// @Description Retrieves the schedule based on department, education form and group number
// @ID get-students-schedule
// @Param education_form path string true "Education form, e.g. `do`"
// @Param department path string true "Department URL, e.g. `knt` for CSIT department"
// @Param group_number path string true "Group number, e.g. `351`"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.StudentsLesson
// @Router /{education_form}/{department}/{group_number} [get]
func (s *server) handleStudentsSchedule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		educationForm := vars["ed_form"]
		departmentUrl := vars["dep_url"]
		groupNum := vars["group_num"]
		l, err := s.store.StudentsSchedule().Select(educationForm, departmentUrl, groupNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if l == nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}
		s.respond(w, http.StatusOK, *l)
	}
}

// @Summary get a list of subgroups of a certain group
// @Tags groups
// @Description Retrieves the subgroups list of a group based on department, education form and group number
// @ID get-group-subgroups
// @Param education_form path string true "Education form, e.g. `do`"
// @Param department path string true "Department URL, e.g. `knt` for CSIT department"
// @Param group_number path string true "Group number, e.g. `351`"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Subgroups
// @Router /{education_form}/{department}/{group_number}/subgroups [get]
func (s *server) handleSubgroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		educationForm := vars["ed_form"]
		departmentUrl := vars["dep_url"]
		groupNum := vars["group_num"]
		l, err := s.store.Subgroups().SelectByGroup(educationForm, departmentUrl, groupNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if l == nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}
		s.respond(w, http.StatusOK, *l)
	}
}
