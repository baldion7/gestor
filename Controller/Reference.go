package Controller

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
)

func GetReferences(c *gin.Context) {
	var references []Model.Reference
	if err := db.ObtenerDB().Preload("Brand").Find(&references).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, references)
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
	if err := c.ShouldBindJSON(&ReferenceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
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
	c.JSON(http.StatusOK, gin.H{"msg": "referencia eliminado exitosamente"})
}

func CreateReference(c *gin.Context) {
	// Obtener los datos del tamanho del cuerpo de la solicitud HTTP
	var ReferenceRequest struct {
		Name              string  `json:"name" binding:"required"`
		BrandId           uint64  `json:"brandId" binding:"required"`
		CostPerProduction float64 `json:"costPerProduction" binding:"required"`
		EnsemblePrice     float64 `json:"ensemblePrice" binding:"required"`
	}

	if err := c.ShouldBindJSON(&ReferenceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la refrencia"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "tamanho creado exitosamente"})
}

func GetReferenceByBrand(c *gin.Context) {
	// Convertir el parámetro BrandId a un entero
	brandId, err := strconv.Atoi(c.Param("brandId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El BrandId proporcionado no es válido"})
		return
	}

	// Buscar las referencias asociadas al BrandId
	var references []Model.Reference
	if err := db.ObtenerDB().Where("brand_id = ?", brandId).Find(&references).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron referencias para el BrandId proporcionado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las referencias"})
		}
		return
	}

	// Devolver las referencias en caso de éxito
	c.JSON(http.StatusOK, references)
}
