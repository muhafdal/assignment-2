package handler

import (
	"h8-assignment-2/docs"
	"h8-assignment-2/infra/config"
	"h8-assignment-2/infra/database"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository/item_pg"
	"h8-assignment-2/repository/order_repository/order_pg"
	"h8-assignment-2/service/order_service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}

func UpdateOrderAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId := ctx.Param("orderId")

		if orderId == "2" {
			forbiddenAccessErr := errs.NewUnauthorizedError("forbidden")
			ctx.AbortWithStatusJSON(forbiddenAccessErr.Status(), forbiddenAccessErr)
			return
		}

		ctx.Next()

	}
}

func DeleteOrderAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId := ctx.Param("orderId")

		if orderId == "2" {

			forbiddenAccessErr := errs.NewUnauthorizedError("forbidden")
			ctx.AbortWithStatusJSON(forbiddenAccessErr.Status(), forbiddenAccessErr)
			return
		}

		ctx.Next()
	}
}

func middleware(ctx *gin.Context) {

}

func StartApp() {
	config.LoadAppConfig()
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewRepository(db)

	itemRepo := item_pg.NewRepository(db)

	orderService := order_service.NewService(orderRepo, itemRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	docs.SwaggerInfo.Title = "H8 Assignment 2"
	docs.SwaggerInfo.Description = "Ini adalah tugas ke 2 dari kelas Kominfo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/orders", orderHandler.CreateOrder)

	r.GET("/orders", Middleware(), orderHandler.GetOrders)

	r.PUT("/orders/:orderId", UpdateOrderAuthorization(), orderHandler.UpdateOrder)

	r.DELETE("/orders/:orderId", DeleteOrderAuthorization(), orderHandler.DeleteOrder)

	r.Run(":8080")
}
