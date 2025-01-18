package Controller

import (
	"fmt"
	"net/http"

	db "gestor/Config/database"
	Model "gestor/Model"
	"github.com/gin-gonic/gin"
)

// CreateSupplier crea un proveedor en la base de datos
func CreateSupplier(c *gin.Context) {
	// Estructura para recibir los datos del proveedor desde el frontend
	var SupplierRequest struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	// Parsear los datos del cuerpo de la solicitud al struct SupplierRequest
	if err := c.ShouldBindJSON(&SupplierRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Imprimir los datos recibidos (para depuración)
	fmt.Println("Datos recibidos:", SupplierRequest)

	// Crear una instancia del modelo de proveedor con los datos recibidos
	supplier := Model.Suppliers{
		Name:    SupplierRequest.Name,
		Contact: SupplierRequest.Contact,
		Email:   SupplierRequest.Email,
		Phone:   SupplierRequest.Phone,
		Address: SupplierRequest.Address,
	}

	// Intentar guardar el proveedor en la base de datos
	if err := db.ObtenerDB().Create(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el proveedor en la base de datos: " + err.Error()})
		return
	}

	// Respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"msg": "Proveedor creado exitosamente", "supplier": supplier})
}

// GetSuppliers obtiene todos los proveedores de la base de datos
func GetSuppliers(c *gin.Context) {
	var suppliers []Model.Suppliers
	if err := db.ObtenerDB().Find(&suppliers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los proveedores"})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID obtiene un proveedor por su ID
func GetSupplierByID(c *gin.Context) {
	var supplier Model.Suppliers
	if err := db.ObtenerDB().First(&supplier, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el proveedor"})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// UpdateSupplier actualiza un proveedor en la base de datos
func UpdateSupplier(c *gin.Context) {
	var supplier Model.Suppliers

	// Buscar el proveedor por ID
	if err := db.ObtenerDB().First(&supplier, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el proveedor"})
		return
	}

	// Estructura para la solicitud de actualización
	var SupplierRequest struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	// Decodificar datos de la solicitud
	if err := c.ShouldBindJSON(&SupplierRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Actualizar los datos del proveedor
	supplier.Name = SupplierRequest.Name
	supplier.Contact = SupplierRequest.Contact
	supplier.Email = SupplierRequest.Email
	supplier.Phone = SupplierRequest.Phone
	supplier.Address = SupplierRequest.Address

	// Guardar los cambios en la base de datos
	if err := db.ObtenerDB().Save(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el proveedor"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// DeleteSupplier elimina un proveedor de la base de datos
func DeleteSupplier(c *gin.Context) {
	var supplier Model.Suppliers

	id := c.Param("id")
	// Eliminar el proveedor
	db.ObtenerDB().Where("id = ?", id).Delete(&supplier)

	c.JSON(http.StatusOK, gin.H{"msg": "Proveedor eliminado exitosamente"})
}
