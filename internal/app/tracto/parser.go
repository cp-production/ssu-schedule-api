package tracto

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "regexp"
)

func GetDepartmentList() ([][]string, error) {
    departmentsList := make([][]string, 0)
    response, err := http.Get(SsuUri)

    if err != nil && response.StatusCode != http.StatusOK {
        return departmentsList, err
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
        return departmentsList, err
    }

    re := regexp.MustCompile(`(?s)<div class="panes_item panes_item__type_group">(.+?)</div>`)
    values := re.FindStringSubmatch(string(body))

    subUL := regexp.MustCompile(`<li><a href='\/schedule\/(.*?)'>(.+?)<\/a><\/li>`)
    liValues := subUL.FindAllStringSubmatch(values[1], -1)

    for _, lv := range liValues {
        href := lv[1]
        facultyName := lv[2]
        departmentsList = append(departmentsList, []string{href, facultyName})
	}
    return departmentsList, err
}

func GetGroupList(departmentID string) ([][]string, error) {
    groupList := make([][]string, 0)
    response, err := http.Get(fmt.Sprintf("%s/%s", SsuUri, departmentID))

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
    reNumOfGroup := regexp.MustCompile(fmt.Sprintf(`%s/(\w+)/(\d+)`, departmentID))
    for _, rows := range values {
        curGroup := reNumOfGroup.FindAllStringSubmatch(rows[0], -1)
        for _, val := range curGroup {
            groupList = append(groupList, []string{val[1], val[2]})
        }
    }

    return groupList, err
}

func GetSchedule(educationForm string, department string, studentGroup string) Schedule {
    response, err := http.Get(fmt.Sprintf("%s/%s/%s/%s",
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
