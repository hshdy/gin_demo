package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"

	"gin_demo/mysql_service"
)

func indexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "这是首页的函数",
	})
}

func someJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO 语言",
		"tag":  "<br>",
	}
	c.AsciiJSON(200, data)
}

func main() {
	fmt.Println("start gin ...")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/index", indexPage)

	r.GET("/someJson", someJson)

	// 引入html页面
	r.LoadHTMLGlob("templates/*")
	r.GET("/indexHtml", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// 和flask 的jinja2 有点像
			"title":   "Main website",
			"taile_1": "22222222",
			"taile_2": "aaaaa",
			"taile_3": "cccccc",
		})
	})

	r.GET("/querySid", mysql_service.QueryGoUserInfoBySid)

	r.Run()

}
