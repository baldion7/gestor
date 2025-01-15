package Controller

import (
	"fmt"
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
)

// GetProducts obtiene todos los productos de la base de datos
func GetProducts(c *gin.Context) {
	var products []Model.Product
	if err := db.ObtenerDB().Preload("Suppliers").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los productos"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID obtiene un producto por su ID
func GetProductByID(c *gin.Context) {
	var product Model.Product
	if err := db.ObtenerDB().First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el producto"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct crea un producto en la base de datos
func CreateProduct(c *gin.Context) {

	// Obtener los datos del producto del cuerpo de la solicitud HTTP
	var ProductRequest struct {
		Name        string  `json:"name" binding:"required"`
		Reference   string  `json:"reference" binding:"required"`
		Color       string  `json:"color" binding:"required"`
		Size        string  `json:"size" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		SuppliersId uint64  `json:"suppliersId" binding:"required"`
		Unitmeasure string  `json:"unitmeasure" binding:"required"`
	}

	// Convertir los datos del request al struct ProductRequest
	if err := c.ShouldBindJSON(&ProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del Modelo de producto con los datos del ProductRequest
	product := Model.Product{
		Name:        ProductRequest.Name,
		Description: ProductRequest.Description,
		Reference:   ProductRequest.Reference,
		Color:       ProductRequest.Color,
		Size:        ProductRequest.Size,
		Price:       ProductRequest.Price,
		SuppliersId: ProductRequest.SuppliersId,
		Unitmeasure: ProductRequest.Unitmeasure,
	}

	// Crea el producto en la base de datos
	if err := db.ObtenerDB().Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el producto"})
		return
	}

	// Devuelve un mensaje de успех
	c.JSON(http.StatusOK, gin.H{"msg": "Producto creado exitosamente"})
}

// UpdateProduct actualiza un producto en la base de datos
func UpdateProduct(c *gin.Context) {
	var product Model.Product

	// Buscar el producto por ID
	if err := db.ObtenerDB().First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el producto"})
		return
	}

	// Estructura para la solicitud de actualización
	var ProductRequest struct {
		Name        string  `json:"name" binding:"required"`
		Reference   string  `json:"reference" binding:"required"`
		Color       string  `json:"color" binding:"required"`
		Size        string  `json:"size" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		SuppliersId uint64  `json:"suppliersId" binding:"required"`
		Unitmeasure string  `json:"unitmeasure" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := c.ShouldBindJSON(&ProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Actualizar los datos del producto
	product.Name = ProductRequest.Name
	product.Description = ProductRequest.Description
	product.Reference = ProductRequest.Reference
	product.Color = ProductRequest.Color
	product.Size = ProductRequest.Size
	product.Price = ProductRequest.Price
	product.SuppliersId = ProductRequest.SuppliersId
	product.Unitmeasure = ProductRequest.Unitmeasure

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto"})
		return
	}

	// Devuelve un mensaje de успех
	c.JSON(http.StatusOK, gin.H{"msg": "Producto actualizado exitosamente"})
}

func DeleteProduct(c *gin.Context) {
	var product Model.Product
	if err := db.ObtenerDB().First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el producto"})
		return
	}
	if err := db.ObtenerDB().Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Producto eliminado exitosamente"})
}

type StockDetail struct {
	Reference    string `json:"reference"`
	Color        string `json:"color"`
	Size         string `json:"size"`
	YardQuantity uint64 `json:"yardQuantity"`
	Stock        int    `json:"stock"`
}

type ProductStock struct {
	Model.Product
	StockDetails []StockDetail    `json:"stockDetails,omitempty"`
	CurrentStock float64          `json:"currentStock"`
	Movements    []Model.Movement `json:"movements"` // Agregar este campo

}

func GetAllProductsStock(c *gin.Context) {
	var products []Model.Product

	// Fetch all products without preload
	if err := db.ObtenerDB().Preload("Movements").Order("id asc").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var stockResults []ProductStock

	for _, product := range products {
		var result ProductStock
		result.Product = product
		result.Movements = product.Movements

		if product.Unitmeasure == "yd" {
			// Para productos medidos en yardas
			var movements []Model.Movement
			if err := db.ObtenerDB().Where("product_id = ?", product.Id).Find(&movements).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Mapa para agrupar por características
			stockByCharacteristics := make(map[string]*StockDetail)

			for _, movement := range movements {
				// Convertir uint64 a string usando strconv
				key := fmt.Sprintf("%s-%s-%s-%d",
					product.Reference,
					product.Color,
					product.Size,
					movement.Quantity)

				if _, exists := stockByCharacteristics[key]; !exists {
					stockByCharacteristics[key] = &StockDetail{
						Reference:    product.Reference,
						Color:        product.Color,
						Size:         product.Size,
						YardQuantity: movement.Quantity,
						Stock:        0,
					}
				}

				if movement.Type == "entrada" {
					stockByCharacteristics[key].Stock++
				} else if movement.Type == "salida" {
					stockByCharacteristics[key].Stock--
				}
			}

			// Convertir mapa a slice y filtrar stock positivo
			for _, detail := range stockByCharacteristics {
				if detail.Stock > 0 {
					result.StockDetails = append(result.StockDetails, *detail)
					result.CurrentStock += float64(detail.Stock)
				}
			}

		} else {
			// Para productos que no son medidos en yardas
			var entradas float64
			var salidas float64

			// Calcular entradas
			db.ObtenerDB().Model(&Model.Movement{}).
				Select("COALESCE(SUM(quantity), 0)").
				Where("product_id = ? AND type = ?", product.Id, "entrada").
				Scan(&entradas)

			// Calcular salidas
			db.ObtenerDB().Model(&Model.Movement{}).
				Select("COALESCE(SUM(quantity), 0)").
				Where("product_id = ? AND type = ?", product.Id, "salida").
				Scan(&salidas)

			result.CurrentStock = entradas - salidas
		}

		stockResults = append(stockResults, result)
	}

	c.JSON(http.StatusOK, stockResults)
}
