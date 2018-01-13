package main

import (
	"crypto/sha1"
	"github.com/gin-gonic/gin"
	"io"
	"sort"
	"fmt"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/wx", func(c *gin.Context) {
		token := "qwer1234"
		timestamp := c.Query("timestamp")
		nonce := c.Query("nonce")

		signature := c.Query("signature")
		echostr := c.Query("echostr")
		list := []string{token, timestamp, nonce}
		hash := sha1.New()
		sort.Strings(list)
		for _, v := range list {
			io.WriteString(hash, v)
		}
		sha := fmt.Sprintf("%x", hash.Sum(nil))
		if sha == signature {
			c.String(200, echostr)
		} else {
			c.String(200, "failed")
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":80")
}
