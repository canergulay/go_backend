package data

import (
	httprequester "backend/global/http_requester"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func SearchCourseUdemy(text string) ParsedHttpMapModel {
	var returnObject ParsedHttpMapModel
	body := strings.NewReader("")
	headers, url := getRequestInfo(text)

	resp, err := httprequester.CreateReqestAndDo("GET", url, body, headers)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println("buradayım")

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if err := json.Unmarshal(buf.Bytes(), &returnObject); err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(returnObject)
	// text is nothing but the course name

	return returnObject
}

func getRequestInfo(text string) ([]map[string]string, string) {
	url := "https://www.udemy.com/api-2.0/courses/?search=" + text

	headers := make([]map[string]string, 1)
	headerOne := make(map[string]string)
	headerOne["Authorization"] = os.Getenv("UDEMY_API_KEY")
	headers[0] = headerOne

	return headers, url
}

type ParsedHttpMapModel struct {
	Aggregations []interface{} `json:"aggregations"`
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

/*
func parseResponseBodyToModel(body map[string]interface{}) ParsedHttpMapModel {
	aggregations := body["aggregations"]
	results := body["results"]
	return ParsedHttpMapModel{Aggregations: aggregations.([]map[string]interface{}), Results: results.([]map[string]interface{})}

}*/
