package course_data_source

import (
	httprequester "backend/global/http_requester"
	"backend/global/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func SearchCourseUdemy(text string, locale string) (ParsedHttpSimplifiedMapModel, error) {

	parsedString := utils.StringSpaceConditioner(text, "+")

	var returnObject ParsedHttpMapModel
	body := strings.NewReader("")
	headers, url := getRequestInfo(parsedString, locale)

	resp, err := httprequester.CreateReqestAndDo(http.MethodGet, url, body, headers)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if err := json.Unmarshal(buf.Bytes(), &returnObject); err != nil {
		fmt.Println(err)
		// circumstance that unmarshal is not successfull means response is not compatible with our expected model
		// which also implies that our search text didn't match anything at the Udemy side
		return ParsedHttpSimplifiedMapModel{}, errors.New("unmarshal was not successfull")
	}
	// text is nothing but the course name

	return parsedModelSimplifier(returnObject)
}

func getRequestInfo(text string, locale string) ([]map[string]string, string) {
	url := "https://www.udemy.com/api-2.0/courses/?search=" + text + "&locale=" + locale

	headers := make([]map[string]string, 1)
	headerOne := make(map[string]string)
	headerOne["Authorization"] = os.Getenv("UDEMY_API_KEY")
	headers[0] = headerOne

	return headers, url
}

func parsedModelSimplifier(modelToSimplify ParsedHttpMapModel) (ParsedHttpSimplifiedMapModel, error) {
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

	var returnModel ParsedHttpSimplifiedMapModel
	var error error

	if aggregationIndex >= 0 {
		returnModel = ParsedHttpSimplifiedMapModel{
			Languages: modelToSimplify.Aggregations[aggregationIndex],
			Results:   modelToSimplify.Results}
		error = nil
	} else {
		returnModel = ParsedHttpSimplifiedMapModel{}
		error = errors.New("couldn't find anything")
	}

	return returnModel, error
}

type ParsedHttpSimplifiedMapModel struct {
	Languages Aggregation `json:"languages"`
	Results   []Course    `json:"results"`
}

type ParsedHttpMapModel struct {
	Aggregations []Aggregation `json:"aggregations"`
	Results      []Course      `json:"results"`
}

type Course struct { //if you want your fields to also be in the lowercase , cover all ofthem within quotation marks.
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Headline    string      `json:"headline"`
	Image       string      `json:"image_480x270"` // IF YOU HAVE HYPEN OR UNDERSCORE WITHIN A FIELD, JUST USE QUATATION MARKS TO INVOLVE FIELD PROPERLY
	Url         string      `json:"url"`
	Instructors []Insructor `json:"visible_instructors"` // SAME SHIT HAPPENS HERE, FIELDS WERE EMPTY BEFORE QUATATION MARKS...
}

type Insructor struct {
	Displayname string `json:"display_name"` // AGAIN QUOTATION MARKS!
	Name        string `json:"name"`
	Image       string `json:"image_100x100"`
}

type Aggregation struct {
	Id      string   `json:"id"`
	Options []Option `json:"options"`
	Title   string   `json:"title"`
}
type Option struct {
	Count int    `json:"count"`
	Key   string `json:"key"`
	Title string `json:"title"`
	Value string `json:"value"`
}
