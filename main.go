package main

import (
	"github.com/dennis-hong/hongi-api/kakaobot"
	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
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
			var kakaoReq kakaobot.KakaoReq
			c.BindJSON(&kakaoReq)
			keys := make([]string, 0, len(kakaoReq.UserAction.Params))
			for k := range kakaoReq.UserAction.Params {
				keys = append(keys, k)
				log.Debugf("keys : %s", keys)
			}
			log.Debugf("kakaoReq : %s", kakaoReq)
			ka := kakaobot.KakaoRes{[]kakaobot.KakaoText{{Type: "text", Text: "test"}}}
			c.JSON(200, ka)
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}