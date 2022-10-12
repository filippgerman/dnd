package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var fio = map[string]string{
	"Kuznetsov": "Peter"}

func GetName(c *gin.Context) {

	nameLast := c.Query("Last name")
	nameName, ok := fio[nameLast]

	if !ok {
		c.JSON(404, gin.H{
			nameLast: ""})
		return
	}
	c.JSON(200, gin.H{
		nameLast: nameName,
	})
}
func PostName(c *gin.Context) {
	nameLast := c.Query("Last name")
	nameName := c.Query("First name")
	if len(nameLast) == 0 || len(nameName) == 0 {

		c.JSON(400, gin.H{
			nameLast: nameName,
		})
		return
	}
	if _, ok := fio[nameLast]; ok {
		c.JSON(409, gin.H{
			"message": "FIO already exists",
		})
		return
	}
	fio[nameLast] = nameName
	c.JSON(201, gin.H{
		nameLast: nameName,
	})
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})
	rGroup := r.Group("/name")
	rGroup.GET("", GetName)
	rGroup.POST("", PostName)
	//rGroup.DELETE("",DeleteName)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
