package model

import "time"

type ESDocument struct {
	Title             string    `json:"title"`
	Url               string    `json:"url"`
	Date              time.Time `json:"date"`
	Content           string    `json:"content"`
	DepartmentID      uint      `json:"department_id"`
	SmallDepartmentID []uint    `json:"small_department_id"`
	ProvinceID        uint      `json:"province_id"`
}

type ESResp struct {
	Hits struct {
		Hits []struct {
			//ID     string      `json:"_id"`
			//Index  string      `json:"_index"`
			//Score  interface{} `json:"_score"`
			Source struct {
				//Content      string    `json:"content"`
				Date              time.Time `json:"date"`
				DepartmentID      int       `json:"department_id"`
				ProvinceID        int       `json:"province_id"`
				SmallDepartmentID []int     `json:"small_department_id"`
				Title             string    `json:"title"`
				URL               string    `json:"url"`
			} `json:"_source"`
			Highlight struct {
				Title   []string `json:"title"`
				Content []string `json:"content"`
			} `json:"highlight"`
		} `json:"hits"`
		Total struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}
