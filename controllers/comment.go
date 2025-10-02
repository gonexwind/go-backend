package controllers

import (
	"gonexwind/backend-api/database"
	"gonexwind/backend-api/helpers"
	"gonexwind/backend-api/models"
	"gonexwind/backend-api/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateComment handles creating a new comment
func CreateComment(c *gin.Context) {
	var req structs.CommentCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// user_id is set by AuthMiddleware
	userID := c.MustGet("user_id").(uint)

	comment := models.Comment{
		Content:   req.Content,
		PostID:    req.PostID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create comment",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Comment created",
		Data:    comment,
	})
}

// UpdateComment handles updating an existing comment by ID
func UpdateComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Comment not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.CommentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	comment.Content = req.Content

	if err := database.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update comment",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Comment updated successfully",
		Data: structs.CommentResponse{
			Id:        comment.Id,
			Content:   comment.Content,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: comment.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// DeleteComment handles deleting a comment by ID
func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Comment not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete comment",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Comment deleted successfully",
	})
}
