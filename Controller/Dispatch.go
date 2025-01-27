package Controller

import (
	db "gestor/Config/database"
	model "gestor/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDispatch(c *gin.Context) {
	var dispatchRequest struct {
		BrandId     uint64   `json:"brand_id" binding:"required"`
		TotalBag    uint64   `json:"total_bag" binding:"required"`
		Collect     string   `json:"collect" binding:"required"`
		Delivery    string   `json:"delivery" binding:"required"`
		Boxes       uint64   `json:"boxes" binding:"required"`
		CutOrderIds []uint64 `json:"cut_order_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&dispatchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create dispatch model
	dispatch := &model.Dispatch{
		BrandId:  dispatchRequest.BrandId,
		TotalBag: dispatchRequest.TotalBag,
		Collect:  dispatchRequest.Collect,
		Delivery: dispatchRequest.Delivery,
		Boxes:    dispatchRequest.Boxes,
	}

	// Find and associate CutOrders
	var cutOrders []*model.CutOrder
	if err := db.ObtenerDB().Find(&cutOrders, dispatchRequest.CutOrderIds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find cut orders"})
		return
	}

	dispatch.CutOrders = cutOrders

	// Create dispatch with associated cut orders
	if err := db.ObtenerDB().Create(dispatch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dispatch)
}

func GetDispatches(c *gin.Context) {
	var dispatches []model.Dispatch

	// Obtén la instancia de la base de datos
	db := db.ObtenerDB()

	// Carga las relaciones necesarias
	err := db.Preload("Brand").
		Preload("CutOrders").
		Preload("CutOrders.Colors").
		Preload("CutOrders.Colors.CutSizes").
		Preload("CutOrders.Reference").
		Find(&dispatches).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Error al obtener las entregas",
			"detail": err.Error(),
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, dispatches)
}
