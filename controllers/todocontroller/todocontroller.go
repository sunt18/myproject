package todocontrollers

import (
	// "encoding/json"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"interview/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	
	var Todo []models.Todo

	acitivity_id := c.Query("activity_group_id")

	// id, _ := strconv.Atoi(acitivity_id)

	models.DB.Where("activity_group_id = ?", acitivity_id).Find(&Todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})

}

func Show(c *gin.Context) {
	var Todo models.Todo
	id := c.Param("id")

	str1 := "Todo with ID "
	str2 := " Not Found"
	result := str1 + id + str2

	if err := models.DB.First(&Todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": result,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func Create(c *gin.Context) {

	var Todo models.Todo

	if err := c.ShouldBindJSON(&Todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	Todo.Priority = "very-high"

	models.DB.Create(&Todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func Update(c *gin.Context) {
	var Todo models.Todo

	id := c.Params.ByName("id")

	str1 := "Todo with ID "
	str2 := " Not Found"
	result := str1 + id + str2

	if err := models.DB.First(&Todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": result,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.BindJSON(&Todo)
	models.DB.Save(&Todo)
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})

}

func Delete(c *gin.Context) {

	id := c.Params.ByName("activity_id")

	var input struct{
		Id json.Number
	}
	
	if err := c.ShouldBindJSON(&input); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage":err.Error()})
		return
	}

	activity_id, _ := input.Id.Int64() 
	
	var Todo models.Todo
	d := models.DB.Where("activity_group_id = ?", id).Delete(&Todo,activity_id)
	fmt.Println(d)

	str1 := "Todo with ID "
	str2 := " Not Found"

	acitivityidconv := strconv.Itoa(int(activity_id))

	// Concatenating strings
	// Using + operator
	result := str1 + acitivityidconv+id + str2

	c.JSON(200, gin.H{
		"status":  "Not Found",
		"message": result,
	})

}
