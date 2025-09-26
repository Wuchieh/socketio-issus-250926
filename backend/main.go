package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zishang520/socket.io/servers/socket/v3"
)

func init() {
	if err := initRedis(); err != nil {
		log.Fatal(err)
	}

	initSocket()
}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.Any("/ws/", func() gin.HandlerFunc {
		return gin.WrapH(socketSer.ServeHandler(nil))
	}())

	r.GET("/sockets", func(c *gin.Context) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		resp := gin.H{}
		socketSer.To("test").FetchSockets()(func(sockets []*socket.RemoteSocket, err error) {
			resp["sockets"] = len(sockets)
			if err != nil {
				resp["error"] = err.Error()
			}
			cancel()
		})
		<-ctx.Done()
		c.JSON(200, resp)
	})

	r.Run("0.0.0.0:" + os.Getenv("PORT"))
}
