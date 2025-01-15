package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func GetCutMovements(c *gin.Context) {
	var cutMovements []Model.CutMovements
	if err := db.ObtenerDB().Find(&cutMovements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, cutMovements)
}

func GetCutMovementByID(c *gin.Context) {
	var cutMovement Model.CutMovements
	if err := db.ObtenerDB().First(&cutMovement, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}
	c.JSON(http.StatusOK, cutMovement)
}

func CreateCutMovement(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del movimiento del cuerpo de la solicitud HTTP
	var CutMovementRequest struct {
		MovementId uint64 `json:"productId" binding:"required"`
		CutOrderId uint64 `json:"quantity" binding:"required"`
	}

	// Convertir los datos del request al struct RolRequest
	if err := mapstructure.Decode(requestData, &CutMovementRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del movimiento: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de rol con los datos del rolRequest
	cutMovement := Model.CutMovements{
		MovementId: CutMovementRequest.MovementId,
		CutOrderId: CutMovementRequest.CutOrderId,
	}

	// Crea el movimiento en la base de datos
	if err := db.ObtenerDB().Create(&cutMovement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el movimiento"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Movimiento creado exitosamente"})
}

func DeleteCutMovement(c *gin.Context) {
	var cutMovement Model.CutMovements
	if err := db.ObtenerDB().First(&cutMovement, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el movimiento"})
		return
	}
	if err := db.ObtenerDB().Delete(&cutMovement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el movimiento"})
		return
	}
	c.JSON(http.StatusOK, cutMovement)
}

func UpdateCutMovement(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del movimiento del cuerpo de la solicitud HTTP
	var CutMovementRequest struct {
		MovementId uint64 `json:"productId" binding:"required"`
		CutOrderId uint64 `json:"quantity" binding:"required"`
	}

	// Convertir los datos del request al struct RolRequest
	if err := mapstructure.Decode(requestData, &CutMovementRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del movimiento: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de rol con los datos del rolRequest
	cutMovement := Model.CutMovements{
		MovementId: CutMovementRequest.MovementId,
		CutOrderId: CutMovementRequest.CutOrderId,
	}

	// Crea el movimiento en la base de datos
	if err := db.ObtenerDB().Save(&cutMovement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el movimiento"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Movimiento creado exitosamente"})
}
