package udemy_models

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
