package rest

import "github.com/gin-gonic/gin"

//RunAPI used in main.go
func RunAPI(adress string) {
	r := gin.Default()

	r.GET("/", handleHome)
	r.GET("/ping", handlePingPong)

	r.Run(adress)
}
