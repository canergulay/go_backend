package course_data_source

import (
	httprequester "backend/global/http_requester"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func SearchCourseUdemy(text string) ParsedHttpSimplifiedMapModel {
	var returnObject ParsedHttpMapModel
	body := strings.NewReader("")
	headers, url := getRequestInfo(text)

	resp, err := httprequester.CreateReqestAndDo("GET", url, body, headers)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if err := json.Unmarshal(buf.Bytes(), &returnObject); err != nil {
		panic(err)
	}
	// text is nothing but the course name

	return parsedModelSimplifier(returnObject)
}

func getRequestInfo(text string) ([]map[string]string, string) {
	url := "https://www.udemy.com/api-2.0/courses/?search=" + text

	headers := make([]map[string]string, 1)
	headerOne := make(map[string]string)
	headerOne["Authorization"] = os.Getenv("UDEMY_API_KEY")
	headers[0] = headerOne

	return headers, url
}

func parsedModelSimplifier(modelToSimplify ParsedHttpMapModel) ParsedHttpSimplifiedMapModel {
	// I EXPECTED AGGREGATION INDEX TO BE 3, SO I WILL DEEM IT IN THAT WAY
	// IF NOT, I WILL SIMPLY ITERATE THROUGH AGGREGATIONS AND FIND THE AGGREGATION WITH THE ID 3

	aggregation := modelToSimplify.Aggregations[3]
	var aggregationIndex int

	if aggregation.Id == "language" {
		aggregationIndex = 3
	} else {
		for index, agg := range modelToSimplify.Aggregations {
			if agg.Id == "language" {
				aggregationIndex = index
				break
			}

		}
	}
	return ParsedHttpSimplifiedMapModel{
		Languages: modelToSimplify.Aggregations[aggregationIndex],
		Results:   modelToSimplify.Results}
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
