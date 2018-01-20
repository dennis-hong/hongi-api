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

type KakaoRes struct {
	Contents []KakaoText `json:"contents"`
}

type KakaoText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/somePost", func(c *gin.Context) {
			var kakaoReq KakaoReq
			c.BindJSON(&kakaoReq)
			keys := make([]string, 0, len(kakaoReq.UserAction.Params))
			for k := range kakaoReq.UserAction.Params {
				keys = append(keys, k)
				log.Println(keys)
			}

			ka := KakaoRes{[]KakaoText{{Type:"text", Text:"test"}}}
			c.JSON(200, ka)
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}