package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCarvings(c *gin.Context) {
	var carvings []Model.Carvings
	if err := db.ObtenerDB().Find(&carvings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los carvings"})
		return
	}
	c.JSON(http.StatusOK, carvings)
}

func GetCarvingByID(c *gin.Context) {
	var carving Model.Carvings
	if err := db.ObtenerDB().First(&carving, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el carvings"})
		return
	}
	c.JSON(http.StatusOK, carving)
}

func CreateCarving(c *gin.Context) {

	// Obtener los datos del carvings del cuerpo de la solicitud HTTP
	var CarvingRequest struct {
		Name               string  `json:"name" binding:"required"`
		Contact            string  `json:"contact" binding:"required"`
		Email              string  `json:"email" binding:"required"`
		Phone              string  `json:"phone" binding:"required"`
		Address            string  `json:"address"`
		ProductionCapacity uint64  `json:"productionCapacity" binding:"required"`
		Delivery           float64 `json:"delivery"`
	}

	// Convertir los datos del request al struct CarvingRequest
	if err := c.ShouldBindJSON(&CarvingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de carvings con los datos del CarvingRequest
	carving := Model.Carvings{
		Name:               CarvingRequest.Name,
		Contact:            CarvingRequest.Contact,
		Email:              CarvingRequest.Email,
		Phone:              CarvingRequest.Phone,
		Address:            CarvingRequest.Address,
		ProductionCapacity: CarvingRequest.ProductionCapacity,
		Delivery:           CarvingRequest.Delivery,
	}

	// Crea el carvings en la base de datos
	if err := db.ObtenerDB().Create(&carving).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el carvings"})
		return
	}
	c.JSON(http.StatusOK, carving)
}

func UpdateCarving(c *gin.Context) {
	var id = c.Param("id")
	// Obtener los datos del carvings del cuerpo de la solicitud HTTP
	var CarvingRequest struct {
		Name               string  `json:"name" binding:"required"`
		Contact            string  `json:"contact" binding:"required"`
		Email              string  `json:"email" binding:"required"`
		Phone              string  `json:"phone" binding:"required"`
		Address            string  `json:"address" binding:"required"`
		ProductionCapacity uint64  `json:"productionCapacity" binding:"required"`
		Delivery           float64 `json:"delivery" `
	}

	// Convertir los datos del request al struct CarvingRequest
	if err := c.ShouldBindJSON(&CarvingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de carvings con los datos del CarvingRequest
	carving := Model.Carvings{
		Name:               CarvingRequest.Name,
		Contact:            CarvingRequest.Contact,
		Email:              CarvingRequest.Email,
		Phone:              CarvingRequest.Phone,
		Address:            CarvingRequest.Address,
		ProductionCapacity: CarvingRequest.ProductionCapacity,
		Delivery:           CarvingRequest.Delivery,
	}

	// Actualiza el carvings en la base de datos
	if err := db.ObtenerDB().Model(&Model.Carvings{}).Where("id = ?", id).Updates(carving).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el carvings"})
		return
	}
	c.JSON(http.StatusOK, carving)

}

func DeleteCarving(c *gin.Context) {
	var carving Model.Carvings
	if err := db.ObtenerDB().First(&carving, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el carvings"})
		return
	}
	if err := db.ObtenerDB().Delete(&carving).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el carvings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Carvings eliminado correctamente"})
}
