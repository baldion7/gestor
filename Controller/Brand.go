package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBrands(c *gin.Context) {
	var brands []Model.Brand
	if err := db.ObtenerDB().Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las marcas"})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func GetBrandByID(c *gin.Context) {
	var brand Model.Brand
	if err := db.ObtenerDB().First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la marca"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	var brand Model.Brand
	if err := db.ObtenerDB().First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la marca"})
		return
	}
	if err := db.ObtenerDB().Delete(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la marca"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Marca eliminada exitosamente"})
}

func CreateBrand(c *gin.Context) {

	// Obtener los datos de la marca del cuerpo de la solicitud HTTP
	var brandRequest struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	// Convertir los datos del request al struct brandRequest
	if err := c.ShouldBindJSON(&brandRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de marca con los datos de brandRequest
	brand := Model.Brand{
		Name:    brandRequest.Name,
		Contact: brandRequest.Contact,
		Email:   brandRequest.Email,
		Phone:   brandRequest.Phone,
		Address: brandRequest.Address,
	}

	// Crea la marca en la base de datos
	if err := db.ObtenerDB().Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la marca"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Marca creada exitosamente"})
}

func UpdateBrand(c *gin.Context) {
	var id = c.Param("id")

	// Obtener los datos de la marca del cuerpo de la solicitud HTTP
	var brandRequest struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	// Convertir los datos del request al struct brandRequest
	if err := c.ShouldBindJSON(&brandRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Crea una instancia del modelo de marca con los datos de brandRequest
	brand := Model.Brand{
		Name:    brandRequest.Name,
		Contact: brandRequest.Contact,
		Email:   brandRequest.Email,
		Phone:   brandRequest.Phone,
		Address: brandRequest.Address,
	}

	// Actualiza la marca en la base de datos
	if err := db.ObtenerDB().Model(&Model.Brand{}).Where("id = ?", id).Updates(brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la marca"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Marca actualizada exitosamente"})
}
