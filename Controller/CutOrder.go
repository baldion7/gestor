package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func GetCutOrderByID(c *gin.Context) {
	var cutSize Model.CutOrder
	if err := db.ObtenerDB().First(&cutSize, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}
	c.JSON(http.StatusOK, cutSize)
}

func GetCutOrders(c *gin.Context) {
	var cutOrders []Model.CutOrder
	if err := db.ObtenerDB().Find(&cutOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, cutOrders)
}

func CreateCutOrder(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var CutOrderRequest struct {
		CreateBy      string  `json:"createBy" binding:"required"`
		Quality       bool    `json:"quality" binding:"required"`
		Arrival       bool    `json:"arrival" binding:"required"`
		Delivered     bool    `json:"delivered" binding:"required"`
		TotalPieces   uint64  `json:"totalPieces" binding:"required"`
		PricePerPiece float64 `json:"totalQuantity" binding:"required"`
		TotalPrice    float64 `json:"totalPrice" binding:"required"`
		Observations  string  `json:"observations" binding:"required"`
		ReferenceId   uint64  `json:"referenceId" binding:"required"`
		CarvingsId    uint64  `json:"carvingsId"`
	}

	// Convertir los datos del request al struct CutOrderRequest
	if err := mapstructure.Decode(requestData, &CutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos de la talla: " + err.Error()})
	}

	// Crea una instancia del modelo de tamanho con los datos del CutOrderRequest
	cutOrder := Model.CutOrder{
		CreateBy:      CutOrderRequest.CreateBy,
		Quality:       CutOrderRequest.Quality,
		Arrival:       CutOrderRequest.Arrival,
		Delivered:     CutOrderRequest.Delivered,
		TotalPieces:   CutOrderRequest.TotalPieces,
		PricePerPiece: CutOrderRequest.PricePerPiece,
		TotalPrice:    CutOrderRequest.TotalPrice,
		Observations:  CutOrderRequest.Observations,
		ReferenceId:   CutOrderRequest.ReferenceId,
		CarvingsId:    CutOrderRequest.CarvingsId,
	}

	// Crea el tamanho en la base de datos
	if err := db.ObtenerDB().Create(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamanho"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "talla creado exitosamente"})
}

func UpdateCutOrder(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var CutOrderRequest struct {
		CreateBy      string  `json:"createBy" binding:"required"`
		Quality       bool    `json:"quality" binding:"required"`
		Arrival       bool    `json:"arrival" binding:"required"`
		Delivered     bool    `json:"delivered" binding:"required"`
		TotalPieces   uint64  `json:"totalPieces" binding:"required"`
		PricePerPiece float64 `json:"totalQuantity" binding:"required"`
		TotalPrice    float64 `json:"totalPrice" binding:"required"`
		Observations  string  `json:"observations" binding:"required"`
		ReferenceId   uint64  `json:"referenceId" binding:"required"`
		CarvingsId    uint64  `json:"carvingsId"`
	}

	// Convertir los datos del request al struct CutOrderRequest
	if err := mapstructure.Decode(requestData, &CutOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos de la talla: " + err.Error()})
	}

	// Buscar el tamanho por ID
	var cutOrder Model.CutOrder
	if err := db.ObtenerDB().First(&cutOrder, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener a talla"})
		return
	}

	// Actualizar los datos del tamanho
	cutOrder.CreateBy = CutOrderRequest.CreateBy
	cutOrder.Quality = CutOrderRequest.Quality
	cutOrder.Arrival = CutOrderRequest.Arrival
	cutOrder.Delivered = CutOrderRequest.Delivered
	cutOrder.TotalPieces = CutOrderRequest.TotalPieces
	cutOrder.PricePerPiece = CutOrderRequest.PricePerPiece
	cutOrder.TotalPrice = CutOrderRequest.TotalPrice
	cutOrder.Observations = CutOrderRequest.Observations
	cutOrder.ReferenceId = CutOrderRequest.ReferenceId
	cutOrder.CarvingsId = CutOrderRequest.CarvingsId

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&cutOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "talla actualizado exitosamente"})
}
