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
	queryFmt = `
  "query": {
    "bool": {
      "must": [
        {
          "bool": {
            "should": [
              {
                "match": {
                  "title": "%s"
                }
              },
              {
                "match": {
                  "content": "%s"
                }
              }
            ]
          }
        } %s
      ]
    }
  },
`
	ResultFmtDSL = `
	{
  %s
  "sort": [
    {
      "date": {
        "order": "asc"
      }
    }
  ],
  "from": %d,
  "size": %d,
  "highlight": %s
}
`

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

	// 分数搜索DSL
	fmtScoreQuery = `
	{
	  "query": {
	    "function_score": {
	      %s
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

func IndexDoc(date time.Time, departmentID, provinceID uint, title, url, content string, spIDs []uint) {
	doc := model.ESDocument{
		Title:             title,
		Url:               url,
		Date:              date,
		Content:           content,
		DepartmentID:      departmentID,
		SmallDepartmentID: spIDs,
		ProvinceID:        provinceID,
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

// SearchDocBySmallDepartmentID 根据小部门筛
func SearchDocBySmallDepartmentID(searchQuery SearchInput, smallDepartmentID, provinceID, from, size int) *model.ESResp {
	var query string
	searchFmt := queryFmtPrint(searchQuery.Text, smallDepartmentID, provinceID)
	if !searchQuery.UseScore {
		query = fmt.Sprintf(ResultFmtDSL,
			searchFmt,
			from-1,
			size,
			highlight)
	} else {
		query = fmt.Sprintf(fmtScoreQuery,
			searchFmt,
			fmtScoreFilters(searchQuery.ScoreField),
			from-1,
			size,
			highlight)
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

func queryFmtPrint(text string, smallDepartmentID, provinceID int) string {
	query := ","
	if smallDepartmentID == 0 && provinceID == 0 {
		query = ""
	} else if smallDepartmentID != 0 && provinceID != 0 {
		query += fmt.Sprintf(`{ "match": { "small_department_id": %d }},`, smallDepartmentID)
		query += fmt.Sprintf(`{ "match": { "province_id": %d }}`, provinceID)
	} else if smallDepartmentID != 0 && provinceID == 0 {
		query += fmt.Sprintf(`{ "match": { "small_department_id": %d }}`, smallDepartmentID)
	} else if smallDepartmentID == 0 && provinceID != 0 {
		query += fmt.Sprintf(`{ "match": { "province_id": %d }}`, provinceID)
	}

	return fmt.Sprintf(queryFmt, text, text, query)
}

func fmtScoreFilters(m map[string]float64) string {
	var filters []string
	for word, weight := range m {
		filter := fmt.Sprintf(fmtScoreFilter, word, word, weight)
		filters = append(filters, filter)
	}
	return strings.Join(filters, ", ")
}
