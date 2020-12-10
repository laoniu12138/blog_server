package v1

import (
	"github.com/gin-gonic/gin"
)

// Tag ...
type Tag struct{}

// NewTag ...
func NewTag() Tag {
	return Tag{}
}

// Get ...
func (t Tag) Get(c *gin.Context) {}

// List ...
func (t Tag) List(c *gin.Context) {}

// Create ...
func (t Tag) Create(c *gin.Context) {}

// Update ...
func (t Tag) Update(c *gin.Context) {}

// Delete ...
func (t Tag) Delete(c *gin.Context) {}
