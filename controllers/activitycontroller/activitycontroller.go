package activitycontrollers

import (
	// "encoding/json"

	"fmt"
	"net/http"

	"interview/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var activity []models.Activity

	models.DB.Find(&activity)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})

}

func Show(c *gin.Context) {
	var Activity models.Activity
	id := c.Param("id")

	str1 := "Activity with ID "
	str2 := " Not Found"
	result := str1 + id + str2

	if err := models.DB.First(&Activity, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": result})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func Create(c *gin.Context) {

	var Activity models.Activity

	if err := c.ShouldBindJSON(&Activity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&Activity)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func Update(c *gin.Context) {
	var Activity models.Activity

	id := c.Params.ByName("id")

	if err := models.DB.First(&Activity, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.BindJSON(&Activity)
	models.DB.Save(&Activity)
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})

}

func Delete(c *gin.Context) {

	id := c.Params.ByName("id")
	var activity models.Activity
	d := models.DB.Where("id = ?", id).Delete(&activity)
	fmt.Println(d)

	str1 := "Activity with ID "
	str2 := " Not Found"
	result := str1 + id + str2

	c.JSON(200, gin.H{
		"status":  "Not Found",
		"message": result,
	})

}
