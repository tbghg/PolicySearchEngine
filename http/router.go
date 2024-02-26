package http

import (
	"PolicySearchEngine/dao/es"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	size = 10
)

func Router() {
	// 创建一个默认的Gin引擎
	router := gin.Default()

	// GET请求处理程序
	router.GET("/search", searchHandel)

	// 启动HTTP服务器，监听端口8080
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("router.Run Failed, err:%+v\n", err)
	}
}

func searchHandel(c *gin.Context) {
	// 获取查询参数
	searchQuery := c.Query("s")

	sDepartmentID := c.Query("did")
	sProvinceID := c.Query("pid")
	sPage := c.Query("page")

	page, err1 := strconv.Atoi(sPage)
	departmentID, err2 := strconv.Atoi(sDepartmentID)
	provinceID, err3 := strconv.Atoi(sProvinceID)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Get请求参数错误\n")
		return
	}

	// 检查查询参数是否为空
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少search参数"})
		return
	}

	esResp := es.SearchDoc(searchQuery, departmentID, provinceID, page, size)
	//totalPage := math.Ceil(float64(esResp.Hits.Total.Value) / size)

	c.JSON(http.StatusOK, esResp.Hits)
}
