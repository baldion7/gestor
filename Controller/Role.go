package Controller

import (
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// CreateRol crea un tamaño en la base de datos
func CreateRole(c *gin.Context) {
	var requestData map[string]interface{}

	// Obtener los datos del tamaño del cuerpo de la solicitud HTTP
	var RolRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	// Convertir los datos del request al struct RolRequest
	if err := mapstructure.Decode(requestData, &RolRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del rol: " + err.Error()})
		return
	}

	// Crea una instancia del modelo de tamaño con los datos del RolRequest
	role := Model.Role{
		Name:        RolRequest.Name,
		Description: RolRequest.Description,
	}

	// Crea el tamaño en la base de datos
	if err := db.ObtenerDB().Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tamaño"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Tamaño creado exitosamente"})
}

// GetRole obtiene todos los tamaños de la base de datos
func GetRole(c *gin.Context) {
	var role []Model.Role
	if err := db.ObtenerDB().Find(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tamaños"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// GetRolByID obtiene un tamaño por su ID
func GetRoleByID(c *gin.Context) {
	var role Model.Role
	if err := db.ObtenerDB().First(&role, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamaño"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// / UpdateRol actualiza un tamaño en la base de datos
func UpdateRole(c *gin.Context) {
	var requestData map[string]interface{}
	var role Model.Role

	// Buscar el tamaño por ID
	if err := db.ObtenerDB().First(&role, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tamaño"})
		return
	}

	// Estructura para la solicitud de actualización
	var RolRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := mapstructure.Decode(requestData, &RolRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del tamaño: " + err.Error()})
		return
	}

	// Actualizar los datos del tamaño
	role.Name = RolRequest.Name
	role.Description = RolRequest.Description

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la talla"})
		return
	}

	c.JSON(http.StatusOK, role)
}
