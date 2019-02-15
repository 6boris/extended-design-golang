package routes

import (
	"extended-design-golang/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) *gin.Engine {

	app.GET("orders", controllers.GetAllOrders)
	app.GET("orders/:order_num", controllers.GetOneOrder)
	app.POST("orders", controllers.CreateOrder)

	app.GET("pay_maps", controllers.GetAllPayMaps)

	return app
}
