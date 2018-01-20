package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type KakaoReq struct {
	UserAction Action `json:"action"`
}

type Action struct {
	Params map[string]string `json:"params"`
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/somePost", func(c *gin.Context) {
		var kakaoReq KakaoReq
		c.BindJSON(kakaoReq)
		keys := make([]string, 0, len(kakaoReq.UserAction.Params))
		for k := range kakaoReq.UserAction.Params {
			keys = append(keys, k)
			log.Println(keys)
		}

		c.JSON(200, gin.H{"user": "", "value": ""})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}