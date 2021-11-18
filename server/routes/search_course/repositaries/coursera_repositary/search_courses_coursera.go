package coursera_repositary

import (
	httprequester "backend/global/http_requester"
	"backend/server/routes/search_course/models/coursera_models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func SearchCourseraCourses(text string) (coursera_models.CourseraResponseModel, error) {
	url := udemyRestApiURLCreator(text)

	var returnObject coursera_models.CourseraResponseModel

	resp, err := httprequester.CreateReqestAndDo(http.MethodGet, url, nil, nil)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err)
		return coursera_models.CourseraResponseModel{}, errors.New("unexpectep error from the api")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if err := json.Unmarshal(buf.Bytes(), &returnObject); err != nil {
		fmt.Println(err)
		// circumstance that unmarshal is not successfull means response is not compatible with our expected model
		// which also implies that our search text didn't match anything at the Udemy side
		return coursera_models.CourseraResponseModel{}, errors.New("unmarshal was not successfull")
	}

	return returnObject, nil

}

func udemyRestApiURLCreator(query string) string {
	u, _ := url.Parse(apiURL)
	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("q", "search")
	values.Set("query", query)
	values.Set("includes", "instructorIds,partnerIds,description,photoUrl")
	values.Set("fields", "instructorIds,partnerIds,description,photoUrl")
	u.RawQuery = values.Encode()
	return u.String()
}

const apiURL = "https://api.coursera.org/api/courses.v1"
