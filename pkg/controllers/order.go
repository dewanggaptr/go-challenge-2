package controllers

import (
	"net/http"

	"gorm.io/gorm"

	"challenge-api/pkg/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (o *OrderController) Index(c *gin.Context) {
	var orders []models.Order

	o.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (o *OrderController) Show(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := o.DB.First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"orderr": order})
}

func (o *OrderController) Create(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	o.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (o *OrderController) Update(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if o.DB.Model(&order).Where("id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func (o *OrderController) Delete(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	// var input struct {
	// 	Id json.Number
	// }

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// // id, _ := input.Id.Int64()
	// if o.DB.Delete(&order, id).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus order"})
	// 	return
	// }

	result := o.DB.First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	result = o.DB.Delete(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
