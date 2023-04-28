package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func main() {
	r := gin.Default()

	r.GET("/get", getIp)
	r.POST("/set", setIp)
	slog.Info("Go is best language!")
	r.Run() // listen and serve on 0.0.0.0:8080
}
