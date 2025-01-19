package Controller

import (
	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var LoginRequest struct {
		Name     string `json:"usuario" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}
	var user Model.User
	if err := db.ObtenerDB().Preload("Role").Where("name = ?", LoginRequest.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequest.Password)); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Inicio de sesión exitoso", "name": user.Name, "id": user.Id, "role": user.Role.Name})
}
