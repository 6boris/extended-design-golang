package boostrap

import (
	"extended-design-golang/app/routes"
	"extended-design-golang/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetApp() *http.Server {

	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(gin.DebugMode)

	app = routes.SetupRouter(app)

	//	初始化依赖
	services.Setup()

	server := &http.Server{
		//Addr:           config.ServerConfig.HttpAddr + ":" + strconv.Itoa(config.ServerConfig.HttpPort),
		Addr:           "127.0.0.1:9000",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}
