package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Likes int `json:"likes"`
	Content string `json:"content"`
	DatePublished string `json:"datePublished"`
	UserID int `json:"userID" gorm:"foreignKey"`
}

func CreateBlog(c *gin.Context) {
	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

fun GetBlog(c *gin.Context) {
  id := c.Param("id")
  var blog Blog

  if err := db.First(&blog, id).Error; err != nil {
  		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
  }

  c.JSON(http.StatusOK, blog)
}

type ListBlogResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	UserID        int    `json:"userId"`
	DatePublished string `json:"datePublished"`
}

func ListBlogs(c *gin.Context) {
	var blogs []BlogResponse

	if err := db.Model(&Blog{}).Select("id, title, user_id, date_published").Scan(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the blogs in the response
	c.JSON(http.StatusOK, blogs)
}



