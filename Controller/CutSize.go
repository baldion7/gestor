package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func GetCutSizeByID(c *gin.Context) {
	var cutSize Model.CutSize
	if err := db.ObtenerDB().First(&cutSize, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}
	c.JSON(http.StatusOK, cutSize)
}

func GetCutSizes(c *gin.Context) {
	var cutSizes []Model.CutSize
	if err := db.ObtenerDB().Find(&cutSizes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, cutSizes)
}

func UpdateCutSize(c *gin.Context) {
	var requestData map[string]interface{}
	var cutSize Model.CutSize

	// Buscar el tamanho por ID
	if err := db.ObtenerDB().First(&cutSize, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener a talla"})
		return
	}

	// Estructura para la solicitud de actualización
	var CutSizeRequest struct {
		Size            string `json:"size" binding:"required"`
		Quantity        uint64 `json:"quantity" binding:"required"`
		ArrivalQuantity uint64 `json:"arrivalQuantity" binding:"required"`
		ColorId         uint64 `json:"colorId" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := mapstructure.Decode(requestData, &CutSizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos de la talla: " + err.Error()})
		return
	}

	// Actualizar los datos del tamanho
	cutSize.Size = CutSizeRequest.Size
	cutSize.Quantity = CutSizeRequest.Quantity
	cutSize.ArrivalQuantity = CutSizeRequest.ArrivalQuantity
	cutSize.ColorId = CutSizeRequest.ColorId

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&cutSize).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}
	c.JSON(http.StatusOK, cutSize)
}

func CreateCutSize(c *gin.Context) {
	var cutSize Model.CutSize
	var requestData map[string]interface{}

	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var CutSizeRequest struct {
		Size            string `json:"size" binding:"required"`
		Quantity        uint64 `json:"quantity" binding:"required"`
		ArrivalQuantity uint64 `json:"arrivalQuantity" binding:"required"`
		ColorId         uint64 `json:"colorId" binding:"required"`
	}

	// Convertir los datos del request al struct CutSizeRequest
	if err := mapstructure.Decode(requestData, &CutSizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos de la talla: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de tamanho con los datos del CutSizeRequest
	cutSize.Size = CutSizeRequest.Size
	cutSize.Quantity = CutSizeRequest.Quantity
	cutSize.ArrivalQuantity = CutSizeRequest.ArrivalQuantity
	cutSize.ColorId = CutSizeRequest.ColorId

	// Crea el tamanho en la base de datos
	if err := db.ObtenerDB().Create(&cutSize).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamanho"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "talla creado exitosamente"})
}
