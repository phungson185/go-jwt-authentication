package dtos

type Time struct {
	Day   int    `json:"day"`
	Month int    `json:"month"`
	Year  int    `json:"year"`
	Week  int    `json:"week"`
	From  string `json:"from"`
	To    string `json:"to"`
	Type_ string `json:"type_"`
}
