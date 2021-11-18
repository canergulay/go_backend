package search_models

type SearchCourseBody struct {
	Source int    `json:"source"` // 0 representing Udemy & 1 representing coursera
	Text   string `json:"text"`   // whatever input is searched by the user
	Locale string `json:"locale"`
}
