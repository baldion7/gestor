package Controller

import (
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func GetReferences(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamaño del cuerpo de la solicitud HTTP
	var ReferenceRequest struct {
		Name              string  `json:"name" binding:"required"`
		BrandId           uint64  `json:"brandId" binding:"required"`
		CostPerProduction float64 `json:"costPerProduction" binding:"required"`
		EnsemblePrice     float64 `json:"ensemblePrice" binding:"required"`
		Description       string  `json:"description" binding:"required"`
	}

	// Convertir los datos del request al struct RolRequest
	if err := mapstructure.Decode(requestData, &ReferenceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del rol: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de tamaño con los datos del RolRequest
	role := Model.Reference{
		Name:              ReferenceRequest.Name,
		BrandId:           ReferenceRequest.BrandId,
		CostPerProduction: ReferenceRequest.CostPerProduction,
		EnsemblePrice:     ReferenceRequest.EnsemblePrice,
	}

	// Crea el tamaño en la base de datos
	if err := db.ObtenerDB().Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamaño"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "referencia creado exitosamente"})
}

func GetReferenceByID(c *gin.Context) {
	var reference Model.Reference
	if err := db.ObtenerDB().First(&reference, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}
	c.JSON(http.StatusOK, reference)
}

func UpdateReference(c *gin.Context) {
	var requestData map[string]interface{}
	var reference Model.Reference

	// Buscar el tamanho por ID
	if err := db.ObtenerDB().First(&reference, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}

	// Estructura para la solicitud de actualización
	var ReferenceRequest struct {
		Name              string  `json:"name" binding:"required"`
		BrandId           uint64  `json:"brandId" binding:"required"`
		CostPerProduction float64 `json:"costPerProduction" binding:"required"`
		EnsemblePrice     float64 `json:"ensemblePrice" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := mapstructure.Decode(requestData, &ReferenceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamanho: " + err.Error()})
		return
	}

	// Actualizar los datos del tamanho
	reference.Name = ReferenceRequest.Name
	reference.BrandId = ReferenceRequest.BrandId
	reference.CostPerProduction = ReferenceRequest.CostPerProduction
	reference.EnsemblePrice = ReferenceRequest.EnsemblePrice

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&reference).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "referencia actualizado exitosamente"})
}

func DeleteReference(c *gin.Context) {
	var reference Model.Reference
	if err := db.ObtenerDB().First(&reference, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamanho"})
		return
	}
	if err := db.ObtenerDB().Delete(&reference).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el tamanho"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Tamanho eliminado exitosamente"})
}

func CreateReference(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var ReferenceRequest struct {
		Name              string  `json:"name" binding:"required"`
		BrandId           uint64  `json:"brandId" binding:"required"`
		CostPerProduction float64 `json:"costPerProduction" binding:"required"`
		EnsemblePrice     float64 `json:"ensemblePrice" binding:"required"`
	}

	// Convertir los datos del request al struct ReferenceRequest
	if err := mapstructure.Decode(requestData, &ReferenceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamanho: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de tamanho con los datos del ReferenceRequest
	reference := Model.Reference{
		Name:              ReferenceRequest.Name,
		BrandId:           ReferenceRequest.BrandId,
		CostPerProduction: ReferenceRequest.CostPerProduction,
		EnsemblePrice:     ReferenceRequest.EnsemblePrice,
	}

	// Crea el tamanho en la base de datos
	if err := db.ObtenerDB().Create(&reference).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamanho"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "tamanho creado exitosamente"})
}
