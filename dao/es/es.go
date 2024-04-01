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

type SearchInput struct {
	Text       string
	UseScore   bool
	ScoreField map[string]float64
}

var (
	es    *elasticsearch.Client
	index string
)

const (
	// 构建高亮查询
	highlight = `
	{
		"fields": {
			"title": {},
			"content": {}
		},
		"fragment_size": 50,
		"pre_tags": ["<em style='color:red'>"],
		"post_tags": ["</em>"]
	}`

	// 一般搜索DSL
	fmtQuery = `
	{
		"query": {
			"bool": {
				"should": [
					%s
				]
			}
		},
		"sort": [
			{ "date": { "order": "asc" }}
		],
		"from": %d,
		"size": %d,
		"highlight": %s
	}`

	// 分数搜索DSL
	fmtScoreQuery = `
	{
	  "query": {
	    "function_score": {
	      "query": {
            "bool": {
              "must": [
	            %s
	          ]
            }
	      },
	      "functions": [
			%s
	      ],
	      "score_mode": "multiply",
	      "boost_mode": "multiply"
	    }
	  },
		"from": %d,
		"size": %d,
		"highlight": %s
	}`
	fmtScoreFilter = `
	{
	  "filter": {
	    "bool": {
	      "should": [
	        { "match": { "title": "%s" }},
	        { "match": { "content": "%s" }}
	      ]
	    }
	  },
	  "weight": %f
	}`

	fmtMustDsl = `
            {
              "bool": {
                "should": [
                  { "match": { "title": "%s" }},
                  { "match": { "content": "%s" }}
                ],
                "minimum_should_match": 1
              }
            }`
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

func SearchDoc(searchQuery SearchInput, departmentID, provinceID, from, size int) *model.ESResp {

	var mustQueries []string
	if searchQuery.UseScore {
		arr := strings.Split(searchQuery.Text, " ")
		for _, s := range arr {
			mustQueries = append(mustQueries,
				fmt.Sprintf(fmtMustDsl, s, s),
			)
		}
	} else {
		mustQueries = append(mustQueries,
			fmt.Sprintf(`{ "match": { "title": "%s" }}`, searchQuery.Text),
			fmt.Sprintf(`{ "match": { "content": "%s" }}`, searchQuery.Text),
		)
	}
	if departmentID != 0 {
		mustQueries = append(mustQueries, fmt.Sprintf(`{ "term": { "department_id": %d }}`, departmentID))
	}
	if provinceID != 0 {
		mustQueries = append(mustQueries, fmt.Sprintf(`{ "term": { "province_id": %d }}`, provinceID))
	}

	var query string
	if searchQuery.UseScore {
		query = fmt.Sprintf(fmtScoreQuery,
			strings.Join(mustQueries, ","),
			fmtScoreFilters(searchQuery.ScoreField),
			from-1,
			size,
			highlight,
		)
	} else {
		query = fmt.Sprintf(fmtQuery,
			strings.Join(mustQueries, ","),
			from-1,
			size,
			highlight,
		)
	}

	fmt.Println(query)

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

func fmtScoreFilters(m map[string]float64) string {
	var filters []string
	for word, weight := range m {
		filter := fmt.Sprintf(fmtScoreFilter, word, word, weight)
		filters = append(filters, filter)
	}
	return strings.Join(filters, ", ")
}
