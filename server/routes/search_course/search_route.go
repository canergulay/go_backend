package searchcourse

import (
	"backend/server/routes/search_course/models/search_models"
	"backend/server/routes/search_course/repositaries/coursera_repositary"
	"backend/server/routes/search_course/repositaries/udemy_repositary"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SearchCourse(c *gin.Context) {
	var body search_models.SearchCourseBody
	c.BindJSON(&body)

	fmt.Println(body, "body burada")
	if body.Source == 0 {
		result, err := udemy_repositary.SearchCourseUdemy(body.Text, body.Locale)
		//if the marshalization process in a erroneus stage,
		//we'll simply return 404, which means searched text couldn't match with anythings
		if err != nil {
			c.JSON(404, errorMessage)
		} else {
			c.JSON(200, result)
		}
	} else {
		result, err := coursera_repositary.SearchCourseraCourses(body.Text)
		if err != nil {
			c.JSON(404, errorMessage)
		} else {
			c.JSON(200, result)
		}
	}
}

const errorMessage = "your search result was empty"
