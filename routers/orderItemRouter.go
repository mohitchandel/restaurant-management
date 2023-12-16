package routes

import (
	controller "restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItem_id", controller.UpdateOrderItem())
}
