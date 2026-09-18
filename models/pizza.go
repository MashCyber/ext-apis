package models

type Pizza struct {
	Code     string `json:"code"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Raw      string `json:"raw"`
	Category string `json:"category"`
}
