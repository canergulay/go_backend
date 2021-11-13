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
		result := course_data_source.SearchCourseUdemy(body.Text)

		c.JSON(200, result)
	}
}
