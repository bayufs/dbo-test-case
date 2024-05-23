package main

import (
	"dbo-test-case/app/routers"
	"fmt"

	"dbo-test-case/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {

	defer config.DisconnectDB(db)
	r := gin.Default()
	routers.InitRouter(r)
	r.Run(fmt.Sprintf("%s:%s", config.SERVICE_HOST, config.SERVICE_PORT))
}
