package model

import "time"

type ESDocument struct {
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	Date         time.Time `json:"date"`
	Content      string    `json:"content"`
	DepartmentID uint      `json:"department_id"`
	ProvinceID   uint      `json:"province_id"`
}

type ESResp struct {
	//Shards struct {
	//	Failed     int `json:"failed"`
	//	Skipped    int `json:"skipped"`
	//	Successful int `json:"successful"`
	//	Total      int `json:"total"`
	//} `json:"_shards"`
	Hits struct {
		Hits []struct {
			//ID     string      `json:"_id"`
			//Index  string      `json:"_index"`
			//Score  interface{} `json:"_score"`
			Source struct {
				Content      string    `json:"content"`
				Date         time.Time `json:"date"`
				DepartmentID int       `json:"department_id"`
				ProvinceID   int       `json:"province_id"`
				Title        string    `json:"title"`
				URL          string    `json:"url"`
			} `json:"_source"`
			//Sort []int64 `json:"sort"`
		} `json:"hits"`
		//MaxScore interface{} `json:"max_score"`
		Total struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}
