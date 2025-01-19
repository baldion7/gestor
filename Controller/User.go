package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateUser crea un usuario en la base de datos
func CreateUser(c *gin.Context) {
	// Obtener los datos del usuario del cuerpo de la solicitud HTTP
	var UserRequest struct {
		Name     string `json:"name" binding:"required"`
		Rol      uint64 `json:"rol" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Convertir los datos del request al struct UserRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	// Encripta la contraseña;
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar la contraseña"})
		return
	}

	// Crea una instancia del modelo de usuario con los datos del UserRequest
	user := Model.User{
		Name:     UserRequest.Name,
		RoleId:   UserRequest.Rol,
		Password: string(hashedPassword),
	}

	// Crea el usuario en la base de datos
	if err := db.ObtenerDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Devuelve un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Usuario creado exitosamente"})
}

// GetUsers obtiene todos los usuarios de la base de datos
func GetUsers(c *gin.Context) {
	var users []Model.User
	if err := db.ObtenerDB().Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser obtiene un usuario de la base de datos
func GetUserByID(c *gin.Context) {
	var user Model.User
	if err := db.ObtenerDB().Preload("Role").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser actualiza un usuario en la base de datos
func UpdateUser(c *gin.Context) {
	var id = c.Param("id")
	var UserRequest struct {
		Name     string `json:"name" binding:"required"`
		Rol      uint64 `json:"rol" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	var user Model.User
	if err := db.ObtenerDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	user.Name = UserRequest.Name
	user.RoleId = UserRequest.Rol
	user.Password = UserRequest.Password
	if err := db.ObtenerDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Usuario actualizado exitosamente"})
}

// DeleteUser elimina un usuario de la base de datos
func DeleteUser(c *gin.Context) {
	var user Model.User
	if err := db.ObtenerDB().First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	if err := db.ObtenerDB().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Usuario eliminado exitosamente"})
}
