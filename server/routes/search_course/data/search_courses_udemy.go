package data

import (
	httprequester "backend/global/http_requester"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func SearchCourseUdemy(text string) map[string]interface{} {
	body := strings.NewReader("")
	headers, url := getRequestInfo(text)

	resp, err := httprequester.CreateReqestAndDo("GET", url, body, headers)

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("buradayım")
	// text is nothing but the course name

	return result
}

func getRequestInfo(text string) ([]map[string]string, string) {
	url := "https://www.udemy.com/api-2.0/courses/?search=" + text

	headers := make([]map[string]string, 1)
	headerOne := make(map[string]string)
	headerOne["Authorization"] = "Basic TkNQczRDMUpBMGFRamlUOEhlRWJtbDdWT3hpMlhKRE5CRnczc292UzpGbU5nVG9PUzBmbzFyZ1Z4dW5DN3l3TmFmMVZlVHdGU2JYeWNnZzNPVFA3QUk4R3d6WlZmdmEwQjVpcDkzQ24wV1p1YlpmTUlGMzc2TzdmZ3dLRXVjMGZkUXNORGNZOXN5ZVh0YkpIYXhzc2s3MXpqeXJCbWUyTWNtSHQ0YVcyNA==A"
	headers[0] = headerOne

	return headers, url
}

type ParsedHttpMapModel struct {
	aggregations string   `json:aggregations`
	results      []Course `json:results`
}

type Course struct {
	title       string      `json:title`
	headline    string      `json:headline`
	image       string      `json:image_240x135`
	url         string      `json:url`
	instructors []Insructor `json:visible_instructors`
}

type Insructor struct {
	displayname string `json:display_name`
	name        string `json:name`
	image       string `json:image_100x100`
}

/*
func parseResponseBodyToModel(body map[string]interface{}) ParsedHttpMapModel {
	aggregations := body["aggregations"]
	results := body["results"]
	return ParsedHttpMapModel{Aggregations: aggregations.([]map[string]interface{}), Results: results.([]map[string]interface{})}

}*/
