package routes

import (
	"github.com/gin-gonic/gin"

	"server/db"
)

func Routes(engine *gin.Engine) {
	engine.GET("/records", GetHistory)
}

func GetHistory(c *gin.Context) {
	db.GetRecords()
}
