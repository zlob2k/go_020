package model

import "time"

type Tdbparam struct {
	DBuser string
	DBpasw string
	DBaddr string
	DBtype string // postgres
	DBname string // links
}
type TNote struct {
	Short_code int
	Url        string
	Created_at time.Time
	Visits     int
}

type TUrlObj struct {
	Url string `json:"url"`
}
