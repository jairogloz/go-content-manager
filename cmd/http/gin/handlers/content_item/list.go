package content_item_hdlr

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) List(c *gin.Context) {

	// Retrieve the 'page' query parameter
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	// Convert the 'page' query parameter to an integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	userID := "1"

	contentItems, err := h.Service.List(userID, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, contentItems)

}
