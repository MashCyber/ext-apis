package models

import "time"

type Norris struct {
	Created_At  string    `json:"created_at"`
	Date_Called time.Time `json:"date_called"`
	URL         string    `json:"url"`
	Value       string    `json:"value"`
}
