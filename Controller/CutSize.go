package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
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

	// Convertir los datos del request al struct CutSizeRequest
	if err := c.ShouldBindJSON(&CutSizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamanho: " + err.Error()})
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

	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var CutSizeRequest struct {
		Size     string `json:"size" binding:"required"`
		Quantity uint64 `json:"quantity" binding:"required"`
		ColorId  uint64 `json:"colorId" binding:"required"`
	}

	// Convertir los datos del request al struct CutSizeRequest
	if err := c.ShouldBindJSON(&CutSizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamanho: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de tamanho con los datos del CutSizeRequest
	cutSize.Size = CutSizeRequest.Size
	cutSize.Quantity = CutSizeRequest.Quantity
	cutSize.ColorId = CutSizeRequest.ColorId
	cutSize.ArrivalQuantity = 0

	// Crea el tamanho en la base de datos
	if err := db.ObtenerDB().Create(&cutSize).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "talla creado exitosamente"})
}

func DeleteCutSizeForColor(c *gin.Context) {
	// Obtener el ColorId del parámetro de la URL
	colorId := c.Param("id")

	// Verificar si el ColorId es válido (no vacío)
	if colorId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de color no proporcionado"})
		return
	}

	// Eliminar todos los tamaños que coincidan con el ColorId
	if err := db.ObtenerDB().Where("color_id = ?", colorId).Delete(&Model.CutSize{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Tallas eliminadas exitosamente"})
}

func UpdateCutSizeArrivalQuantity(c *gin.Context) {
	var cutSize Model.CutSize

	// Buscar el tamanho por ID
	if err := db.ObtenerDB().First(&cutSize, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener a talla"})
		return
	}

	// Estructura para la solicitud de actualización
	var CutSizeRequest struct {
		ArrivalQuantity uint64 `json:"arrivalQuantity" binding:"required"`
	}

	// Convertir los datos del request al struct CutSizeRequest
	if err := c.ShouldBindJSON(&CutSizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamanho: " + err.Error()})
		return
	}

	// Actualizar los datos del tamanho
	cutSize.ArrivalQuantity = CutSizeRequest.ArrivalQuantity

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&cutSize).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}
	c.JSON(http.StatusOK, cutSize)
}
