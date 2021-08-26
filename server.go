package main

import (
	"github.com/gin-gonic/gin"
	menuItemsController "store.tn/api/src/menu-items/controller"
	metaController "store.tn/api/src/meta/controller"
	productController "store.tn/api/src/products/controller"
	"store.tn/api/src/shared/database"
)

func main() {
	router := gin.Default()
	router.GET("/products", productController.GetAll)
	router.POST("/products", productController.CreateProduct)
	router.GET("/products/:id", productController.GetById)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.GET("/products/sub-category/:subCategoryId", productController.GetBySubCategoryId)
	router.POST("/products/:productId/sub-category/:subCategoryId", productController.AddProductToSubCategory)
	router.DELETE("/products/:productId/sub-category/:subCategoryId", productController.RemoveProductFromSubCategory)

	router.GET("/meta/menu-items", menuItemsController.GetMenuItems)
	router.POST("/meta/menu-items", menuItemsController.CreateMenuItem)
	router.GET("/meta/meta-data", metaController.GetMetaData)

	router.Run("localhost:8080")

	ctx := database.GetContext(30)
	defer database.Client.Disconnect(ctx)
}
