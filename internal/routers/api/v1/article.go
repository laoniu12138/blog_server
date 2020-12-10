package v1

import (
	"github.com/gin-gonic/gin"
)

// Article ...
type Article struct{}

// NewArticle ...
func NewArticle() Tag {
	return Tag{}
}

// Get ...
func (t Article) Get(c *gin.Context) {}

// List ...
func (t Article) List(c *gin.Context) {}

// Create ...
func (t Article) Create(c *gin.Context) {}

// Update ...
func (t Article) Update(c *gin.Context) {}

// Delete ...
func (t Article) Delete(c *gin.Context) {}
