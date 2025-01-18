package Controller

import (
	"fmt"
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
)

func GetMovements(c *gin.Context) {
	var movements []Model.Movement
	if err := db.ObtenerDB().Find(&movements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los movimientos"})
		return
	}
	c.JSON(http.StatusOK, movements)
}

func DeleteMovement(c *gin.Context) {
	var movement Model.Movement
	if err := db.ObtenerDB().First(&movement, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el movimiento"})
		return
	}
	if err := db.ObtenerDB().Delete(&movement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el movimiento"})
		return
	}
	c.JSON(http.StatusOK, movement)
}

func CreateMovement(c *gin.Context) {
	var movement Model.Movement
	fmt.Print(c.Request.Body)

	// Obtener los datos del movimiento del cuerpo de la solicitud HTTP
	var MovementRequest struct {
		Type      string `json:"type" binding:"required"`
		ProductId uint64 `json:"ProductId" binding:"required"`
		Reason    string `json:"reason" binding:"required"`
		Quantity  uint64 `json:"quantity" binding:"required"`
	}

	// Convertir los datos del request al struct RolRequest
	if err := c.ShouldBindJSON(&MovementRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de rol con los datos del rolRequest
	movement.Type = MovementRequest.Type
	movement.ProductId = MovementRequest.ProductId
	movement.Quantity = MovementRequest.Quantity
	movement.Reason = MovementRequest.Reason

	// Crea el movimiento en la base de datos
	if err := db.ObtenerDB().Create(&movement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el movimiento"})
		return
	}

	// Devuelve el movimiento creado
	c.JSON(http.StatusOK, movement)
}

func UpdateMovement(c *gin.Context) {
	var movement Model.Movement

	var MovementRequest struct {
		Type      string `json:"type" binding:"required"`
		ProductId uint64 `json:"productId" binding:"required"`
		Quantity  uint64 `json:"quantity" binding:"required"`
		Reason    string `json:"reason" binding:"required"`
	}

	if err := c.ShouldBindJSON(&MovementRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := db.ObtenerDB().First(&movement, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el movimiento"})
		return
	}

	movement.Type = MovementRequest.Type
	movement.ProductId = MovementRequest.ProductId
	movement.Quantity = MovementRequest.Quantity
	movement.Reason = MovementRequest.Reason

	if err := db.ObtenerDB().Save(&movement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el movimiento"})
		return
	}
}
