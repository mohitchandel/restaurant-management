package main

import (
	"os"
	"restaurant-management/database"
	"restaurant-management/middleware"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "nil" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)
	routes.NotesRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.UserRoutes(router)

	router.Run(":", port)

}
