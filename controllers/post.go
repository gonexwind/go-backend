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

func CreatePost(c *gin.Context) {
	var req = structs.PostCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Auth middleware sets user_id as uint in context
	userID := c.MustGet("user_id").(uint)

	post := models.Post{
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
		OwnerID:   userID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create post",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Post created",
		Data:    post,
	})

}

func ShowPosts(c *gin.Context) {
	var posts []models.Post

	// Include comments for each post
	if err := database.DB.Preload("Comments").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to fetch posts",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Posts",
		Data:    posts,
	})
}

func UpdatePost(c *gin.Context) {

	// Ambil ID user dari parameter URL
	id := c.Param("id")

	// Inisialisasi post
	var post models.Post

	// Cari post berdasarkan ID
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Post not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	//struct post request
	var req = structs.PostUpdateRequest{}

	// Bind JSON request ke struct UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Update post dengan data baru
	post.Title = req.Title
	post.Body = req.Body

	// Simpan perubahan ke database
	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update post",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
		Data: structs.PostResponse{
			Id:        post.Id,
			Title:     post.Title,
			Body:      post.Body,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeletePost(c *gin.Context) {

	// Ambil ID user dari parameter URL
	id := c.Param("id")

	// Inisialisasi post
	var post models.Post

	// Cari post berdasarkan ID
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Post not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Hapus post dari database
	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete post",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Post deleted successfully",
	})
}

// ShowPost returns a single post with its comments
func ShowPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := database.DB.Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Post not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Post Detail",
		Data:    post,
	})
}
