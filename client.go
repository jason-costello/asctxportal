package asctxportal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	districtID string
	CSRFToken  string
	students   []StudentSummary
}

func (c *Client) GetBaseURL() *url.URL{
	return c.baseURL
}

func (c *Client) GetHttpClient() *http.Client{
	return c.httpClient
}

func (c *Client) GetStudent() ([]Student, error) {

	return nil, nil

}

func New(u *url.URL, districtID string) *Client {

	jar, _ := cookiejar.New(nil)

	hc := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       10 * time.Second,
	}

	return &Client{
		httpClient: hc,
		baseURL:    u,
		districtID: districtID,
	}
}

func (c *Client) Login(user, pass string) ([]StudentSummary, error) {

	// grab initial page to get csrf token and cookies prior to posting login attempt

	lr, err := c.httpClient.Get(fmt.Sprintf("%s/login?distid=%s", c.baseURL, c.districtID))
	if err != nil {
		return nil, err
	}

	loginPage, err := ioutil.ReadAll(lr.Body)
	if err != nil {
		log.Fatal(err)
	}
	csrfToken := getCSRFToken(string(loginPage))
	lr.Body.Close()

	c.CSRFToken = csrfToken

	params := url.Values{}
	params.Set("username", user)
	params.Set("password", pass)
	params.Set("signin", ",")
	params.Set("_csrf", csrfToken)
	payload := bytes.NewBufferString(params.Encode())
	loginRequest, err := http.NewRequest("POST", c.baseURL.String()+"/loginPP", payload)
	if err != nil {
		panic(err)
	}

	// Headers
	loginRequest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	loginRequest.Header.Add("Origin", "https://asctxportal.esc13.net")
	loginRequest.Header.Add("Accept-Encoding", "gzip, deflate, gzip, deflate, br")
	loginRequest.Header.Add("Host", "asctxportal.esc13.net")
	loginRequest.Header.Add("Accept-Language", "en-us")
	loginRequest.Header.Add("Referer", "https://asctxportal.esc13.net/ParentPortal/login?distid=046901")
	loginRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	loginResponse, err := c.httpClient.Do(loginRequest)
	if err != nil {
		panic(err)
	}

	defer loginResponse.Body.Close()
	if loginResponse.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", loginResponse.StatusCode, loginResponse.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(loginResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Grab students information for the landing page just after logging in.
	var snames []string
	var sids []string
	var schoolNames []string
	var galerts []string
	var aalerts []string

	// Find the review items
	doc.Find("div.studentPanelList").Each(func(i int, s *goquery.Selection) {

		s.Find("div.stupanel").Each(func(i int, si *goquery.Selection) {

			if v, ok := si.Attr("data-name"); ok {
				snames = append(snames, v)
			}
			if v, ok := si.Attr("data-id"); ok {
				sids = append(sids, v)
			}

		})

		var cnt = 0
		var idx = 0
		s.Find("p.text-right").Each(func(x int, students *goquery.Selection) {

			if cnt%3 == 0 && (cnt != 0 && idx != 0) {
				idx++
			}

			sParts := strings.Split(students.Text(), ":")
			if len(sParts) == 0 {
				panic("failed")
			} else if len(sParts) < 2 && len(sParts) > 0 {
				// school  name found
				schoolNames = append(schoolNames, sParts[0])
			} else if len(sParts) == 2 {

				alertType := strings.ToLower(strings.TrimSpace(sParts[0]))

				if alertType == "attendance alerts" {
					aalerts = append(aalerts, strings.TrimSpace(sParts[1]))
				} else if alertType == "grade alerts" {
					galerts = append(galerts, strings.TrimSpace(sParts[1]))

				}

			}

			cnt++

		})

	})

	var students []StudentSummary
	for i, s := range snames {

		ga, err := strconv.Atoi(galerts[i])
		if err != nil {
			ga = 0
		}
		aa, err := strconv.Atoi(aalerts[i])
		if err != nil {
			aa = 0
		}
		students = append(students, StudentSummary{
			ID:               sids[i],
			Name:             strings.ReplaceAll(s, "\t", ""),
			SchoolName:       schoolNames[i],
			GradeAlerts:      ga,
			AttendanceAlerts: aa,
		})
	}

	return students, nil
}

type StudentSummary struct {
	ID               string
	Name             string
	SchoolName       string
	GradeAlerts      int
	AttendanceAlerts int
}

func getCSRFToken(page string) string {

	re := regexp.MustCompile(`<meta name="_csrf" content="(?P<csrf>[a-zA-Z0-9-]{1,})" />`)
	matches := re.FindStringSubmatch(page)
	csrfIndex := re.SubexpIndex("csrf")
	csrfToken := matches[csrfIndex]
	return csrfToken
}
