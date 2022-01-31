package routes

import (
	"github.com/abdulaleem-git/restaurant-mgmt-system/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
