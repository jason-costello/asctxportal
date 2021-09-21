package asctxportal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
type : "GET",
contentType : "application/json",
url : "/ParentPortal/assignments/findAssignments",
data : {
pCourseID : $("#course").val(),
pCycle : $("#cycle").val(),
view : $("#view").val(),
pMaxDueDate : $("#duedate").val().replace(/\//g,""),
pIncludeBlankDueDates : document.getElementById("blanksearch").checked,
csrfmiddlewaretoken: $("#csrfmiddlewaretoken").val()
},
dataType : 'json',
timeout : 100000,
success : function(data) {
if(data.code == "1"){
var returnData = {};
returnData.data = data.data;
console.log(returnData)
*/

func UnmarshalAssignments(data []byte) (Assignments, error) {
	var r Assignments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Assignments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Assignments struct {
	Code    int64   `json:"code"`
	Message string  `json:"message"`
	Data    []Datum `json:"data"`
}

type Datum struct {
	Assignment     string   `json:"assignment"`
	AssignmentNote string   `json:"assignmentNote"`
	Category       Category `json:"category"`
	Course         string   `json:"course"`
	DueDate        string   `json:"dueDate"`
	FailingGrade   bool     `json:"failingGrade"`
	Grade          string   `json:"grade"`
}

type Category string

const (
	DailyGrades    Category = "Daily Grades"
	HomeworkGrades Category = "Homework Grades"
	MajorGrades    Category = "Major Grades"
)

type GetGradeOpts map[string]string

func GetGrades(opts GetGradeOpts, c *Client) (Assignments, error) {

	c.baseURL.Path = "/ParentPortal/assignments/findAssignments"

	q := c.baseURL.Query()

	for k, v := range opts {
		q.Set(k, v)
	}

	c.baseURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", c.baseURL.String(), nil)
	fmt.Println(req.URL.String())
	if err != nil {
		return Assignments{}, err
	}
	req.Header.Set("contentType", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Assignments{}, err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Assignments{}, err
	}
	return UnmarshalAssignments(d)

}

func SetStudent(studentID string, c *Client) error {

	c.baseURL.Path = "/ParentPortal/assignments/selectStudent"
	q := c.baseURL.Query()

	q.Set("pCourseID", "All")
	q.Set("studentId", studentID)
	q.Set("csrfmiddlewaretoken", c.CSRFToken)
	q.Set("pCycle", "All")

	c.baseURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", c.baseURL.String(), nil)
	fmt.Println(req.URL.String())
	if err != nil {
		return err
	}
	req.Header.Set("contentType", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil

}
