package main

import (
	"assignment2/docs"
	_ "assignment2/docs"
	"assignment2/handler"
	"assignment2/infra"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

var appPort = os.Getenv("PORT")

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := infra.RouterInit()
	db := infra.MySqlInit()

	order := handler.NewOrderHandler(db)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("order", order.CreateOrder)
	r.GET("order/:order_id", order.GetOrder)
	r.PUT("order/:order_id", order.UpdateOrder)
	r.DELETE("order/:order_id", order.DeleteOrder)
	log.Fatal(r.Run(":" + appPort))
}
