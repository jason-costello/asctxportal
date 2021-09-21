package asctxportal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UnmarshalAttendance(data []byte) (Attendance, error) {
	var r Attendance
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Attendance) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Attendance struct {
	Campus               Campus                 `json:"Campus"`
	CurrentSemester      int64                  `json:"currentSemester"`
	Code                 int64                  `json:"code"`
	DetailAttendanceList []DetailAttendanceList `json:"detailAttendanceList"`
	Semester             map[string]Semester    `json:"semester"`
	Student              Student                `json:"Student"`
	Message              string                 `json:"message"`
}

type Campus struct {
	CurrentSemester             int64               `json:"currentSemester"`
	CurrentCycle                int64               `json:"currentCycle"`
	SemestersPerYear            int64               `json:"semestersPerYear"`
	CyclesPerSemester           int64               `json:"cyclesPerSemester"`
	AdaPeriod                   int64               `json:"adaPeriod"`
	FailingGrade                int64               `json:"failingGrade"`
	PointScaleLow               int64               `json:"pointScaleLow"`
	PointScaleHigh              int64               `json:"pointScaleHigh"`
	PointGoal                   int64               `json:"pointGoal"`
	AdaAttendance               bool                `json:"adaAttendance"`
	HasExam                     bool                `json:"hasExam"`
	HasCitizenship              bool                `json:"hasCitizenship"`
	DisplayNumericGradesAsAlpha bool                `json:"displayNumericGradesAsAlpha"`
	ID                          string              `json:"id"`
	Name                        string              `json:"name"`
	DefaultCitizenshipGrade     string              `json:"defaultCitizenshipGrade"`
	LowGradeLevel               interface{}         `json:"lowGradeLevel"`
	SemesterRanges              map[string]Semester `json:"semesterRanges"`
	Notes                       interface{}         `json:"notes"`
	DistrictNotes               interface{}         `json:"districtNotes"`
	RawCampusNotes              []interface{}       `json:"rawCampusNotes"`
	RawDistNotes                []interface{}       `json:"rawDistNotes"`
	LowGrdLvl                   string              `json:"lowGrdLvl"`
	CampusNoteFontSize          string              `json:"campusNoteFontSize"`
	GradeReportingOptions       interface{}         `json:"gradeReportingOptions"`
	Demographic                 interface{}         `json:"demographic"`
	GradebookOptions            GradebookOptions    `json:"gradebookOptions"`
	CurrentYear                 interface{}         `json:"currentYear"`
	NumberOfSemesters           int64               `json:"numberOfSemesters"`
	NumberOfCycles              int64               `json:"numberOfCycles"`
}

type GradebookOptions struct {
	AllowPriorAttendancePosting       bool                 `json:"allowPriorAttendancePosting"`
	SpecialProgramsToDisplay          []string             `json:"specialProgramsToDisplay"`
	OverrideSemester                  bool                 `json:"overrideSemester"`
	AddCategoriesLock                 bool                 `json:"addCategoriesLock"`
	PGPFlag                           string               `json:"pgpFlag"`
	DisciplineFlag                    string               `json:"disciplineFlag"`
	GenericProgramsToDisplay          string               `json:"genericProgramsToDisplay"`
	Alloweditgradesforpreviouscycle   bool                 `json:"alloweditgradesforpreviouscycle"`
	TeacherAllowStandardsBasedGrading bool                 `json:"teacherAllowStandardsBasedGrading"`
	AllowStandardsBasedGrading        bool                 `json:"allowStandardsBasedGrading"`
	GradeConversionModel              GradeConversionModel `json:"gradeConversionModel"`
	UsesElementarySkillsReportCard    bool                 `json:"usesElementarySkillsReportCard"`
	AllowWgtTypePercentFlag           string               `json:"allowWgtTypePercentFlag"`
	AllowWgtTypePointsFlag            string               `json:"allowWgtTypePointsFlag"`
	AllowWgtTypeMultiplierFlag        string               `json:"allowWgtTypeMultiplierFlag"`
	AllowRubrics                      bool                 `json:"allowRubrics"`
	ConvertMissingStandardToZero      bool                 `json:"convertMissingStandardToZero"`
	StandardScoreType                 string               `json:"standardScoreType"`
	AllowRptCardNarrative             string               `json:"allowRptCardNarrative"`
	UsesBehaviorGrading               bool                 `json:"usesBehaviorGrading"`
	AttendanceCampus                  bool                 `json:"attendanceCampus"`
	GradebookCampus                   bool                 `json:"gradebookCampus"`
}

type GradeConversionModel struct {
	GradeConversionID int64 `json:"gradeConversionID"`
	PointScaleLow     int64 `json:"pointScaleLow"`
	PointScaleHigh    int64 `json:"pointScaleHigh"`
}

type Semester struct {
	Semester string `json:"semester"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	EndCyc1  string `json:"endCyc1"`
	EndCyc2  string `json:"endCyc2"`
	EndCyc3  string `json:"endCyc3"`
}

type DetailAttendanceList struct {
	StudentID         string           `json:"studentId"`
	SemesterID        interface{}      `json:"semesterId"`
	Period            string           `json:"period"`
	Course            string           `json:"course"`
	Instructor        string           `json:"instructor"`
	DetailCodeList    [][]CodeList     `json:"detailCodeList"`
	SimpleCodeList    [][]CodeList     `json:"simpleCodeList"`
	CodeHeaderList    []CodeHeaderList `json:"codeHeaderList"`
	HasInstructorNote bool             `json:"hasInstructorNote"`
	InstructorNote    interface{}      `json:"instructorNote"`
	InstructorEmail   string           `json:"instructorEmail"`
	Withdrawn         bool             `json:"withdrawn"`
}

type CodeList struct {
	AttendanceID     int64            `json:"attendanceID"`
	StudentID        string           `json:"studentID"`
	CampusID         string           `json:"campusId"`
	Date             int64            `json:"date"`
	AbsDate          string           `json:"absDate"`
	Period           int64            `json:"period"`
	Code             Code             `json:"code"`
	RcardEquivalent  Code             `json:"rcardEquivalent"`
	RcardDescription RcardDescription `json:"rcardDescription"`
	CourseNumber     *string          `json:"courseNumber"`
	AttendanceStyle  AttendanceStyle  `json:"attendanceStyle"`
}

type Student struct {
	StudentID              string       `json:"studentId"`
	StudentRegKey          string       `json:"studentRegKey"`
	FirstName              string       `json:"firstName"`
	MiddleName             string       `json:"middleName"`
	LastName               string       `json:"lastName"`
	Birthdate              int64        `json:"birthdate"`
	Gender                 string       `json:"gender"`
	StudentPicture         string       `json:"studentPicture"`
	UnreadAlert            int64        `json:"unreadAlert"`
	EnrollmentDate         interface{}  `json:"enrollmentDate"`
	IsExcludedFromGrading  bool         `json:"isExcludedFromGrading"`
	Track                  int64        `json:"track"`
	Grade                  string       `json:"grade"`
	AddressLine1           string       `json:"addressLine1"`
	AddressLine2           string       `json:"addressLine2"`
	ContactInfo            interface{}  `json:"contactInfo"`
	PhoneAC                string       `json:"phoneAc"`
	PhoneNum               string       `json:"phoneNum"`
	Enrollments            []Enrollment `json:"enrollments"`
	EmailAssociationStatus int64        `json:"emailAssociationStatus"`
	StudentNameFML         string       `json:"studentNameFML"`
	StudentNameLFM         string       `json:"studentNameLFM"`
}

type Enrollment struct {
	StudentEnrollID  string      `json:"studentEnrollID"`
	CampusID         string      `json:"campusId"`
	StudentID        interface{} `json:"studentId"`
	WithdrawalDate   string      `json:"withdrawalDate"`
	EntryDate        string      `json:"entryDate"`
	EnrollmentStatus int64       `json:"enrollmentStatus"`
	GradeLevel       string      `json:"gradeLevel"`
	Track            string      `json:"track"`
	StrNum           string      `json:"strNum"`
	StrName          string      `json:"strName"`
	AptNum           string      `json:"aptNum"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	Zip              string      `json:"zip"`
	Zip4             string      `json:"zip4"`
}

type CodeHeaderList string

type AttendanceStyle string

type Code string

type RcardDescription string

func GetAttendance(c *Client, semesterID, studentID string) error {

	c.baseURL.Path = "/ParentPortal/attendance/totalattendance/" + studentID

	q := c.baseURL.Query()

	q.Set("semesterId", semesterID)
	q.Set("csrfmiddlewaretoken", c.CSRFToken)

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

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(d))
	return nil

}
