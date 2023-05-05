package tracto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func GetDepartmentList() (Departments, error) {
	response, err := http.Get(fmt.Sprintf("%s/departments", TractoUri))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	departmentsList := Departments{}
	err = json.NewDecoder(response.Body).Decode(&departmentsList)
	if err != nil {
		panic(err)
	}

	return departmentsList, nil

}

func GetGroupList(dep_url string) ([][]string, error) {
	groupList := make([][]string, 0)
	response, err := http.Get(fmt.Sprintf("%s/%s", SsuUri, dep_url))

	if err != nil && response.StatusCode != http.StatusOK {
		return groupList, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return groupList, err
	}

	re := regexp.MustCompile(`(?s)<div class="fieldset-wrapper">(.+?)</div>`)
	values := re.FindAllStringSubmatch(string(body), -1)
	reNumOfGroup := regexp.MustCompile(fmt.Sprintf(`%s/(\w+)/(\d+)`, dep_url))
	for _, rows := range values {
		curGroup := reNumOfGroup.FindAllStringSubmatch(rows[0], -1)
		for _, val := range curGroup {
			groupList = append(groupList, []string{val[1], val[2]})
		}
	}

	return groupList, nil
}

func GetSchedule(educationForm string, department string, studentGroup string) Schedule {
	response, err := http.Get(fmt.Sprintf("%s/schedule/%s/%s/%s",
		TractoUri,
		educationForm,
		department,
		studentGroup))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	schedule := Schedule{}

	err = json.NewDecoder(response.Body).Decode(&schedule)
	if err != nil {
		panic(err)
	}

	return schedule
}
