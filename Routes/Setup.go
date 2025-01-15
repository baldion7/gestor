package Routes

import (
	"gestor/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/api/users", Controller.GetUsers)
	router.GET("/api/user/:id", Controller.GetUserByID)
	router.POST("/api/user", Controller.CreateUser)
	router.PUT("/api/user/:id", Controller.UpdateUser)
	router.DELETE("/api/user/:id", Controller.DeleteUser)

	router.GET("/api/suppliers", Controller.GetSuppliers)
	router.GET("/api/suppliers/:id", Controller.GetSupplierByID)
	router.POST("/api/suppliers", Controller.CreateSupplier)
	router.PUT("/api/suppliers/:id", Controller.UpdateSupplier)
	router.DELETE("/api/suppliers/:id", Controller.DeleteSupplier)

	router.GET("/api/sizes", Controller.GetSizes)
	router.GET("/api/size/:id", Controller.GetSizeByID)
	router.POST("/api/size", Controller.CreateSize)
	router.PUT("/api/size/:id", Controller.UpdateSize)

	router.GET("/api/roles", Controller.GetRole)
	router.GET("/api/role/:id", Controller.GetRoleByID)
	router.POST("/api/role", Controller.CreateRole)
	router.PUT("/api/role/:id", Controller.UpdateRole)

	router.GET("/api/referencia", Controller.GetReferences)
	router.GET("/api/referencia/:id", Controller.GetReferenceByID)
	router.POST("/api/referencia", Controller.CreateReference)
	router.PUT("/api/referencia/:id", Controller.UpdateReference)

	router.GET("/api/product", Controller.GetProducts)
	router.GET("/api/product/:id", Controller.GetProductByID)
	router.POST("/api/product", Controller.CreateProduct)
	router.PUT("/api/product/:id", Controller.UpdateProduct)
	router.DELETE("/api/product/:id", Controller.DeleteProduct)

	router.GET("/api/products-stock", Controller.GetAllProductsStock)

	router.GET("/api/movements", Controller.GetMovements)
	router.POST("/api/movements", Controller.CreateMovement)
	router.PUT("/api/movements/:id", Controller.UpdateMovement)
	router.DELETE("/api/movements/:id", Controller.DeleteMovement)

	router.GET("/api/cutsizes", Controller.GetCutSizes)
	router.GET("/api/cutsizes/:id", Controller.GetCutSizeByID)
	router.POST("/api/cutsizes", Controller.CreateCutSize)
	router.PUT("/api/cutsizes/:id", Controller.UpdateCutSize)

	router.GET("/api/cut-orders", Controller.GetCutOrders)
	router.GET("/api/cut-order/:id", Controller.GetCutOrderByID)
	router.POST("/api/cut-order", Controller.CreateCutOrder)
	router.PUT("/api/cut-order/:id", Controller.UpdateCutOrder)

	router.GET("/api/cut-movement", Controller.GetCutMovements)
	router.GET("/api/cut-movement/:id", Controller.GetCutMovementByID)
	router.POST("/api/cut-movement", Controller.CreateCutMovement)
	router.PUT("/api/cut-movement/:id", Controller.UpdateCutMovement)
	router.DELETE("/api/cut-movement/:id", Controller.DeleteCutMovement)

	router.GET("/api/colors", Controller.GetColors)
	router.GET("/api/color/:id", Controller.GetColorByID)
	router.POST("/api/color", Controller.CreateColor)
	router.PUT("/api/color/:id", Controller.UpdateColor)
	router.DELETE("/api/color/:id", Controller.DeleteColor)

	router.GET("/api/carvings", Controller.GetCarvings)
	router.GET("/api/carving/:id", Controller.GetCarvingByID)
	router.POST("/api/carving", Controller.CreateCarving)
	router.PUT("/api/carving/:id", Controller.UpdateCarving)
	router.DELETE("/api/carving/:id", Controller.DeleteCarving)

	router.GET("/api/brand", Controller.GetBrands)
	router.GET("/api/brand/:id", Controller.GetBrandByID)
	router.POST("/api/brand", Controller.CreateBrand)
	router.PUT("/api/brand/:id", Controller.UpdateBrand)
	router.DELETE("/api/brand/:id", Controller.DeleteBrand)

}
