package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func GetColors(c *gin.Context) {
	var colors []Model.Color
	if err := db.ObtenerDB().Find(&colors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los colores"})
		return
	}
	c.JSON(http.StatusOK, colors)
}

func GetColorByID(c *gin.Context) {
	var color Model.Color
	if err := db.ObtenerDB().First(&color, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el color"})
		return
	}
	c.JSON(http.StatusOK, color)
}

func DeleteColor(c *gin.Context) {
	var color Model.Color
	if err := db.ObtenerDB().First(&color, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el color"})
		return
	}
	if err := db.ObtenerDB().Delete(&color).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el color"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Color eliminado exitosamente"})
}

func CreateColor(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del color del cuerpo de la solicitud HTTP
	var ColorRequest struct {
		Name        string  `json:"name" binding:"required"`
		TotalPieces uint64  `json:"totalPieces" binding:"required"`
		TotalPrice  float64 `json:"totalPrice" binding:"required"`
		CutOrderId  uint64  `json:"cutOrderId" binding:"required"`
	}

	// Convertir los datos del request al struct ColorRequest
	if err := mapstructure.Decode(requestData, &ColorRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del color: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de color con los datos del ColorRequest
	color := Model.Color{
		Name:        ColorRequest.Name,
		TotalPieces: ColorRequest.TotalPieces,
		TotalPrice:  ColorRequest.TotalPrice,
		CutOrderId:  ColorRequest.CutOrderId,
	}

	// Crea el color en la base de datos
	if err := db.ObtenerDB().Create(&color).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el color"})
		return
	}
}

func UpdateColor(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del color del cuerpo de la solicitud HTTP
	var ColorRequest struct {
		Name        string  `json:"name" binding:"required"`
		TotalPieces uint64  `json:"totalPieces" binding:"required"`
		TotalPrice  float64 `json:"totalPrice" binding:"required"`
		CutOrderId  uint64  `json:"cutOrderId" binding:"required"`
	}

	// Convertir los datos del request al struct ColorRequest
	if err := mapstructure.Decode(requestData, &ColorRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del color: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de color con los datos del ColorRequest
	color := Model.Color{
		Name:        ColorRequest.Name,
		TotalPieces: ColorRequest.TotalPieces,
		TotalPrice:  ColorRequest.TotalPrice,
		CutOrderId:  ColorRequest.CutOrderId,
	}

	// Crea el color en la base de datos
	if err := db.ObtenerDB().Save(&color).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el color"})
		return
	}
}
