package routes

import (
	controller "restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menus_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menus_id", controller.UpdateMenu())
}
