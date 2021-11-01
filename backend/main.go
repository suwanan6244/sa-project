package main

import (
	"github.com/suwanan6244/sa-project/controller"

	"github.com/suwanan6244/sa-project/entity"

	"github.com/suwanan6244/sa-project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Staff Routes
			protected.GET("/staffs", controller.ListStaffs)
			protected.GET("/staff/:id", controller.GetStaff)
			protected.PATCH("/staffs", controller.UpdateStaff)
			protected.DELETE("/staffs/:id", controller.DeleteStaff)

			// ProductType Routes
			protected.GET("/producttypes", controller.ListProductTypes)
			protected.GET("/producttype/:id", controller.GetProductType)
			protected.POST("/producttypes", controller.CreateProductType)
			protected.PATCH("/producttypes", controller.UpdateProductType)
			protected.DELETE("/producttypes/:id", controller.DeleteProductType)

			// Product Routes
			protected.GET("/products", controller.ListProducts)
			protected.GET("/product/:id", controller.GetProduct)
			protected.POST("/products", controller.CreateProduct)
			protected.PATCH("/products", controller.UpdateProduct)
			protected.DELETE("/products/:id", controller.DeleteProduct)

			// Supplier Routes
			protected.GET("/suppliers", controller.ListSuppliers)
			protected.GET("/supplier/:id", controller.GetSupplier)
			protected.POST("/suppliers", controller.CreateSupplier)
			protected.PATCH("/suppliers", controller.UpdateSupplier)
			protected.DELETE("/suppliers/:id", controller.DeleteSupplier)

			// ProductStock Routes
			protected.GET("/product_stocks", controller.ListProductStocks)
			protected.GET("/productstock/:id", controller.GetProductStock)
			protected.POST("/product_stocks", controller.CreateProductStock)
			protected.PATCH("/product_stocks", controller.UpdateProductStock)
			protected.DELETE("/productstocks/:id", controller.DeleteProductStock)

		}
	}

	// Staff Routes
	r.POST("/staffs", controller.CreateStaff)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
