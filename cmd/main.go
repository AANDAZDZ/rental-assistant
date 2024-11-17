package main

import (
	"github.com/gin-gonic/gin"
	"rental-assistant/internal/server"
)

func main() {
	engine := gin.Default()
	engine.POST("/getNearRegionInfo", server.GetNearRegionInfo)
	err := engine.Run()
	if err != nil {
		return
	}
}
