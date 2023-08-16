package handlers

import (
	"net/http"
	"notes/pkg/models"
	"notes/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var LastID = 0

func CreateNote(c *gin.Context) {
	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	LastID++

	note.ID = LastID
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	models.Notes = append(models.Notes, note)

	utils.SaveNotesToFile()

	c.JSON(http.StatusCreated, gin.H{
		"id": note.ID,
	})
}

func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, models.Notes)
}

func GetNote(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, val := range models.Notes {
		if val.ID == id {
			c.JSON(http.StatusOK, val)
		}
	}
}

func UpdateNote(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for index := range models.Notes {
		if models.Notes[index].ID == id {
			models.Notes[index].Content = note.Content
			models.Notes[index].UpdatedAt = time.Now()
		}
	}

	utils.SaveNotesToFile()

	c.JSON(http.StatusOK, gin.H{
		"reason": "successfully updated",
	})
}

func DeleteNote(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var requiredIndex int

	for index := range models.Notes {
		if models.Notes[index].ID == id {
			requiredIndex = index
		}
	}

	if requiredIndex < 0 || requiredIndex >= len(models.Notes) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.Notes = append(models.Notes[:requiredIndex], models.Notes[requiredIndex+1:]...)

	utils.SaveNotesToFile()

	c.JSON(http.StatusOK, gin.H{
		"reason": "successfully deleted",
	})
}
