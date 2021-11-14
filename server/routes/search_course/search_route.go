package searchcourse

import (
	"backend/server/routes/search_course/course_data_source"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SearchCourseBody struct {
	Source int    `json:"source"` // 0 representing Udemy & 1 representing coursera
	Text   string `json:"text"`   // whatever input is searched by the user
}

func SearchCourse(c *gin.Context) {
	var body SearchCourseBody
	c.BindJSON(&body)

	fmt.Println(body, "body burada")
	if body.Source == 0 {
		result, err := course_data_source.SearchCourseUdemy(body.Text)
		//if the marshalization process in a erroneus stage,
		//we'll simply return 404, which means searched text couldn't match with anythings
		if err != nil {
			c.JSON(404, errorMessage)
		} else {
			c.JSON(200, result)
		}
	}
}

const errorMessage = "your search result was empty"
