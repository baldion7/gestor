package Controller

import (
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// CreateSize crea un tamaño en la base de datos
func CreateSize(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamaño del cuerpo de la solicitud HTTP
	var SizeRequest struct {
		Name string `json:"name" binding:"required"`
	}

	// Convertir los datos del request al struct SizeRequest
	if err := mapstructure.Decode(requestData, &SizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamaño: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de tamaño con los datos del SizeRequest
	size := Model.Sizes{
		Name: SizeRequest.Name,
	}

	// Crea el tamaño en la base de datos
	if err := db.ObtenerDB().Create(&size).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamaño"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Tamaño creado exitosamente"})
}

// GetSizes obtiene todos los tamaños de la base de datos
func GetSizes(c *gin.Context) {
	var sizes []Model.Sizes
	if err := db.ObtenerDB().Find(&sizes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, sizes)
}

// GetSizeByID obtiene un tamaño por su ID
func GetSizeByID(c *gin.Context) {
	var size Model.Sizes
	if err := db.ObtenerDB().First(&size, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamaño"})
		return
	}
	c.JSON(http.StatusOK, size)
}

// / UpdateSize actualiza un tamaño en la base de datos
func UpdateSize(c *gin.Context) {
	var requestData map[string]interface{}
	var size Model.Sizes

	// Buscar el tamaño por ID
	if err := db.ObtenerDB().First(&size, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamaño"})
		return
	}

	// Estructura para la solicitud de actualización
	var SizeRequest struct {
		Name string `json:"name" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := mapstructure.Decode(requestData, &SizeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamaño: " + err.Error()})
		return
	}

	// Actualizar los datos del tamaño
	size.Name = SizeRequest.Name

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&size).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}

	c.JSON(http.StatusOK, size)
}
