package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test01(c *gin.Context) {
	data := map[string]string{"name": "zhj", "message": "hello world"}
	res := &AmisResp{
		Status: 200,
		Msg:    "ok",
		Data:   data,
	}
	fmt.Println(res)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   data,
	})
}
