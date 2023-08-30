package controllers

import (
	"AvitoProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateSegment_HTTP(c *gin.Context) {
	var input models.Segment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segment, err := (&input).Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"segment": segment})
}

func DeleteSegment_HTTP(c *gin.Context) {
	var input models.Segment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*var toDelete *models.Segment
	models.DB.Where()*/
	segment, err := (&input).Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"segment deleted": segment})
}

func AddUserToSegment_HTTP(c *gin.Context) {
	input := struct {
		gorm.Model
		UserID           uint     `gorm:"not null" json:"user_id"`
		SegmentsToAdd    []string `json:"segments_to_add"`
		SegmentsToDelete []string `json:"segments_to_delete"`
	}{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var segmentsToAdd []models.Segment
	var segmentsToDelete []models.Segment

	if err := models.DB.Where("slug IN (?)", input.SegmentsToAdd).Find(&segmentsToAdd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := models.DB.Where("slug IN (?)", input.SegmentsToDelete).Find(&segmentsToDelete).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := models.AddUserToSegment(segmentsToAdd, segmentsToDelete, input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func ActiveSegments_HTTP(c *gin.Context) {
	input_user_id := struct {
		userId uint `json:"user_id"`
	}{}

	if err := c.ShouldBindJSON(&input_user_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segments, err := models.ActiveSegments(input_user_id.userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var segments_slugs []string
	for _, seg := range segments {
		segments_slugs = append(segments_slugs, seg.Slug)
	}
	fmt.Println(segments_slugs)
	c.JSON(http.StatusOK, gin.H{"active_segments": segments})
}
