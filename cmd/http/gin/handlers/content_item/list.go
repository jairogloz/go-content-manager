package content_item_hdlr

import "github.com/gin-gonic/gin"

func (h *HttpHandler) List(c *gin.Context) {

	userID := "1"

	contentItems, err := h.Service.List(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, contentItems)

}
