package es

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
	"time"
)

var (
	es    *elasticsearch.Client
	index string
)

func Init() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			config.V.GetString("es.addr"),
		},
	}
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("ES初始化错误，err:%+v\n", err.Error())
	}
	index = config.V.GetString("es.index")
}

func IndexDoc(date time.Time, departmentID, provinceID uint, title, url, content string) {
	doc := model.ESDocument{
		Title:        title,
		Url:          url,
		Date:         date,
		Content:      content,
		DepartmentID: departmentID,
		ProvinceID:   provinceID,
	}
	data, _ := json.Marshal(doc)
	idx, err := es.Index(index, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("ES上传数据失败，err:%+v\n", err)
		return
	}
	fmt.Println(idx.String())
}

func MatchAllDoc() {
	query := `{ "query": { "match_all": {} } }`
	search, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return
	}
	fmt.Println(search.String())
}

func SearchDoc(searchQuery string, departmentID, provinceID, from, size int) *model.ESResp {

	var mustQueries []string
	mustQueries = append(mustQueries,
		fmt.Sprintf(`{ "match": { "title": "%s" }}`, searchQuery),
		fmt.Sprintf(`{ "match": { "content": "%s" }}`, searchQuery),
	)
	if departmentID != 0 {
		mustQueries = append(mustQueries, fmt.Sprintf(`{ "term": { "department_id": %d }}`, departmentID))
	}
	if provinceID != 0 {
		mustQueries = append(mustQueries, fmt.Sprintf(`{ "term": { "province_id": %d }}`, provinceID))
	}

	query := `
	{
		"query": {
			"bool": {
				"must": [
					%s
				]
			}
		},
		"sort": [
			{ "date": { "order": "asc" }}
		],
		"from": %d,
		"size": %d
	}`

	query = fmt.Sprintf(query, strings.Join(mustQueries, ","), from, size)

	searchResult, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		fmt.Println("Error executing search:", err)
		return nil
	}

	// 解析 searchResult 中的 JSON 数据
	var responseData model.ESResp
	_ = json.NewDecoder(searchResult.Body).Decode(&responseData)

	return &responseData
}
