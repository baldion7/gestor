package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CutOrder struct {
	ID           uint          `json:"id"`
	Colors       []Color       `json:"colors" gorm:"foreignKey:CutOrderID"`
	Brand        Brand         `json:"brand" gorm:"foreignKey:BrandID"`
	CutMovements []CutMovement `json:"cutMovements" gorm:"foreignKey:CutOrderID"`
}

type Color struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	CutSizes []CutSize `json:"cutSizes" gorm:"foreignKey:ColorID"`
}

type CutSize struct {
	ID   uint   `json:"id"`
	Size string `json:"size"`
}

type Brand struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CutMovement struct {
	ID        uint       `json:"id"`
	Movements []Movement `json:"movements" gorm:"foreignKey:CutMovementID"`
}

type Movement struct {
	ID       uint      `json:"id"`
	Type     string    `json:"type"`
	Products []Product `json:"products" gorm:"foreignKey:MovementID"`
}

type Product struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCutOrderByID(c *gin.Context) {
	var cutOrders []Model.CutOrder

	result := db.ObtenerDB().
		Preload("Colors").
		Preload("Colors.CutSizes").
		Preload("CutMovements").
		Preload("CutMovements.Movement").
		Preload("CutMovements.Movement.Product").
		Preload("Carvings").
		Preload("Reference").
		Preload("Reference.Brand"). // Si necesitas la información de la marca de la referencia
		First(&cutOrders, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener las órdenes de corte",
			"details": result.Error.Error(),
		})
		return
	}

	if len(cutOrders) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron órdenes de corte",
			"data":    []Model.CutOrder{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Órdenes de corte obtenidas exitosamente",
		"data":    cutOrders,
	})
}

// Función para obtener las órdenes de corte
func GetCutOrders(c *gin.Context) {
	var cutOrders []Model.CutOrder

	result := db.ObtenerDB().
		Preload("Colors").
		Preload("Colors.CutSizes").
		Preload("CutMovements").
		Preload("CutMovements.Movement").
		Preload("CutMovements.Movement.Product").
		Preload("Carvings").
		Preload("Reference").
		Preload("Reference.Brand"). // Si necesitas la información de la marca de la referencia
		Find(&cutOrders)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener las órdenes de corte",
			"details": result.Error.Error(),
		})
		return
	}

	if len(cutOrders) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron órdenes de corte",
			"data":    []Model.CutOrder{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Órdenes de corte obtenidas exitosamente",
		"data":    cutOrders,
	})
}

func CreateCutOrder(c *gin.Context) {
	// Estructura para los datos de la solicitud
	var cutOrderRequest struct {
		CreatedBy     string  `json:"createdBy" binding:"required"`
		Observations  string  `json:"observations" binding:"required"` // Changed JSON tag to match field name
		ReferenceId   uint64  `json:"referenceId" binding:"required"`
		Quality       bool    `json:"quality" `
		Delivered     bool    `json:"delivered" `
		TotalPieces   uint64  `json:"totalPieces" binding:"required"`
		PricePerPiece float64 `json:"pricePerPiece" binding:"required"`
		TotalPrice    float64 `json:"totalPrice" binding:"required"`
	}

	// Parsear los datos de la solicitud al struct
	if err := c.ShouldBindJSON(&cutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la instancia del modelo CutOrder
	cutOrder := Model.CutOrder{
		CreateBy:      cutOrderRequest.CreatedBy,
		Observations:  cutOrderRequest.Observations,
		ReferenceId:   cutOrderRequest.ReferenceId,
		Quality:       cutOrderRequest.Quality,
		Delivered:     cutOrderRequest.Delivered,
		TotalPieces:   cutOrderRequest.TotalPieces,
		PricePerPiece: cutOrderRequest.PricePerPiece,
		TotalPrice:    cutOrderRequest.TotalPrice,
	}

	// Guardar la orden en la base de datos
	if err := db.ObtenerDB().Create(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la orden: " + err.Error()})
		return
	}

	// Devolver una respuesta exitosa con los datos creados
	c.JSON(http.StatusCreated, gin.H{
		"message": "Orden de corte creada exitosamente",
		"data":    cutOrder,
	})
}

func UpdateCutOrder(c *gin.Context) {
	// Estructura para los datos de la solicitud
	var cutOrderRequest struct {
		CreatedBy     string  `json:"createdBy" binding:"required"`
		Observations  string  `json:"observations" binding:"required"`
		ReferenceId   uint64  `json:"referenceId" binding:"required"`
		Quality       bool    `json:"quality"`
		Arrival       bool    `json:"arrival"`
		Delivered     bool    `json:"delivered"`
		TotalPieces   uint64  `json:"totalPieces" binding:"required"`
		PricePerPiece float64 `json:"pricePerPiece" binding:"required"`
		TotalPrice    float64 `json:"totalPrice" binding:"required"`
		CarvingsId    *uint64 `json:"carvingsId"` // Campo opcional
	}

	// Parsear los datos de la solicitud al struct
	if err := c.ShouldBindJSON(&cutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar la orden de corte por ID
	var cutOrder Model.CutOrder
	if err := db.ObtenerDB().First(&cutOrder, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden de corte no encontrada"})
		return
	}

	// Preparar el mapa de actualización
	updates := map[string]interface{}{
		"create_by":       cutOrderRequest.CreatedBy,
		"quality":         cutOrderRequest.Quality,
		"arrival":         cutOrderRequest.Arrival,
		"delivered":       cutOrderRequest.Delivered,
		"total_pieces":    cutOrderRequest.TotalPieces,
		"price_per_piece": cutOrderRequest.PricePerPiece,
		"total_price":     cutOrderRequest.TotalPrice,
		"observations":    cutOrderRequest.Observations,
		"reference_id":    cutOrderRequest.ReferenceId,
	}

	// Solo incluir CarvingsId en la actualización si se proporciona en la solicitud
	if cutOrderRequest.CarvingsId != nil {
		updates["carvings_id"] = *cutOrderRequest.CarvingsId
	}

	// Actualizar usando el mapa de campos
	if err := db.ObtenerDB().Model(&cutOrder).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la orden de corte"})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Orden de corte actualizada exitosamente"})
}

func UpdateCutOrderCarving(c *gin.Context) {
	// Estructura para los datos de la solicitud
	var cutOrderRequest struct {
		CarvingsId uint64 `json:"carvingsId" binding:"required"`
	}

	// Parsear los datos de la solicitud al struct
	if err := c.ShouldBindJSON(&cutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar la orden de corte por ID
	var cutOrder Model.CutOrder
	if err := db.ObtenerDB().First(&cutOrder, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden de corte no encontrada"})
		return
	}

	// Actualizar el campo CarvingsId
	cutOrder.CarvingsId = cutOrderRequest.CarvingsId
	if err := db.ObtenerDB().Save(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la orden de corte"})
	}

	// Responder con válido
	c.JSON(http.StatusOK, gin.H{"msg": "CarvingsId actualizado exitosamente"})
}

func UpdateCutOrderObservations(c *gin.Context) {
	var id = c.Param("id")
	var cutOrderRequest struct {
		Observations string `json:"observation" binding:"required"`
	}

	if err := c.ShouldBindJSON(&cutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cutOrder Model.CutOrder
	if err := db.ObtenerDB().First(&cutOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden de corte no encontrada"})
		return
	}

	cutOrder.Observations2 = cutOrderRequest.Observations
	if err := db.ObtenerDB().Save(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la orden de corte"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Observaciones actualizadas exitosamente"})
}

func UpdateCutOrderFinish(c *gin.Context) {
	var id = c.Param("id")
	var cutOrderRequest struct {
		Finish bool `json:"finish" binding:"required"`
	}

	if err := c.ShouldBindJSON(&cutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cutOrder Model.CutOrder
	if err := db.ObtenerDB().First(&cutOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden de corte no encontrada"})
		return
	}

	cutOrder.Finish = cutOrderRequest.Finish
	if err := db.ObtenerDB().Save(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la orden de corte"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Orden de corte actualizada exitosamente"})
}
