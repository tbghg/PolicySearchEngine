package http

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/dao/es"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	size = 10
)

func Router() {
	router := gin.Default()
	router.GET("/search", searchHandel)
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("router.Run Failed, err:%+v\n", err)
	}
}

func searchHandel(c *gin.Context) {

	searchQuery := c.Query("s")
	sDepartmentID := c.Query("did")
	sProvinceID := c.Query("pid")
	sPage := c.Query("page")

	page, err1 := strconv.Atoi(sPage)
	departmentID, err2 := strconv.Atoi(sDepartmentID)
	provinceID, err3 := strconv.Atoi(sProvinceID)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Get请求参数错误\n，err1:%+v，err2:%+v，err3:%+v", err1, err2, err3)
		return
	}

	// 检查查询参数是否为空
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少search参数"})
		return
	}

	esResp := es.SearchDoc(preSearch(searchQuery), departmentID, provinceID, page, size)
	//totalPage := math.Ceil(float64(esResp.Hits.Total.Value) / size)

	c.JSON(http.StatusOK, esResp.Hits)
}

func preSearch(s string) string {

	preSearchUrl := config.V.GetString("http.pre-search-url") + url.QueryEscape(s)
	resp, err := http.Get(preSearchUrl)
	if err != nil {
		fmt.Println("Error occurred while sending GET request:", err)
		return s
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return s
	}

	var r struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return s
	}

	if resp.StatusCode == http.StatusOK && r.Code == 200 {
		return r.Message
	} else {
		fmt.Printf("Non-200 status code: %d\n, resp:%+v", resp.StatusCode, r)
	}
	return s
}
