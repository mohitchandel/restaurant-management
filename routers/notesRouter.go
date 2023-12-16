package routes

import (
	controller "restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controller.GetNotes())
	incomingRoutes.GET("/note/:note_id", controller.GetNote())
	incomingRoutes.POST("/notes", controller.CreateNote())
	incomingRoutes.PATCH("/note/:note_id", controller.UpdateNote())
}
