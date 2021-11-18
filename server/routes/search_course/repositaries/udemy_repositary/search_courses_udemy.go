package udemy_repositary

import (
	httprequester "backend/global/http_requester"
	"backend/global/utils"
	"backend/server/routes/search_course/models/udemy_models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func SearchCourseUdemy(text string, locale string) (udemy_models.ParsedHttpSimplifiedMapModel, error) {

	parsedString := utils.StringSpaceConditioner(text, "+")

	var returnObject udemy_models.ParsedHttpMapModel

	headers, url := getRequestInfo(parsedString, locale)

	resp, err := httprequester.CreateReqestAndDo(http.MethodGet, url, nil, headers)
	if err != nil || resp.StatusCode != 200 {
		log.Fatalln(err)
		return udemy_models.ParsedHttpSimplifiedMapModel{}, errors.New("unexpected error")
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if err := json.Unmarshal(buf.Bytes(), &returnObject); err != nil {
		fmt.Println(err)
		// circumstance that unmarshal is not successfull means response is not compatible with our expected model
		// which also implies that our search text didn't match anything at the Udemy side
		return udemy_models.ParsedHttpSimplifiedMapModel{}, errors.New("unmarshal was not successfull")
	}
	// text is nothing but the course name

	return parsedModelSimplifier(returnObject)
}

func getRequestInfo(text string, locale string) ([]map[string]string, string) {

	apiURL := "https://www.udemy.com/api-2.0/courses/"
	u, _ := url.Parse(apiURL)
	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("search", text)
	values.Set("locale", locale)
	u.RawQuery = values.Encode()

	headers := make([]map[string]string, 1)
	headerOne := make(map[string]string)
	headerOne["Authorization"] = os.Getenv("UDEMY_API_KEY")
	headers[0] = headerOne

	return headers, u.String()
}

func parsedModelSimplifier(modelToSimplify udemy_models.ParsedHttpMapModel) (udemy_models.ParsedHttpSimplifiedMapModel, error) {
	// I EXPECTED AGGREGATION INDEX TO BE 3, SO I WILL DEEM IT IN THAT WAY
	// IF NOT, I WILL SIMPLY ITERATE THROUGH AGGREGATIONS AND FIND THE AGGREGATION WITH THE ID 3
	fmt.Println(modelToSimplify)
	aggregations := modelToSimplify.Aggregations
	aggregationIndex := -1

	if len(aggregations) > 3 && aggregations[3].Id == "language" {

		aggregationIndex = 3
	} else {
		for index, agg := range modelToSimplify.Aggregations {
			if agg.Id == "language" {
				aggregationIndex = index
				break
			}

		}
	}

	var returnModel udemy_models.ParsedHttpSimplifiedMapModel
	var error error

	if aggregationIndex >= 0 {
		returnModel = udemy_models.ParsedHttpSimplifiedMapModel{
			Languages: modelToSimplify.Aggregations[aggregationIndex],
			Results:   modelToSimplify.Results}
		error = nil
	} else {
		returnModel = udemy_models.ParsedHttpSimplifiedMapModel{}
		error = errors.New("couldn't find anything")
	}

	return returnModel, error
}
