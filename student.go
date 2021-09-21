package asctxportal

//func GetAttendance(semesterID int, studentID string, c *Client) (Attendance, error){
//
//	var resp, err = c.httpClient.Get(fmt.Sprintf(`https://asctxportal.esc13.net/ParentPortal/attendance/detailattendance/%s?semesterId=%d`, studentID, semesterID))
//	if err != nil{
//		return Attendance{}, err
//	}
//
//	defer resp.Body.Close()
//	if resp.StatusCode != 200 {
//		return Attendance{}, errors.New(resp.Status)
//	}
//	attendanceData, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return Attendance{}, err
//	}
//
//
//	return UnmarshalAttendance(attendanceData)
//
//}
