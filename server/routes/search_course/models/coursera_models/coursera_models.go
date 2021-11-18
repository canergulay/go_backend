package coursera_models

type CourseraResponseModel struct {
	Courses []CourseraCourse `json:"elements"`
}

type CourseraCourse struct {
	PhotoUrl    string `json:"photoUrl"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
