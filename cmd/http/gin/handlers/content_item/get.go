package content_item_hdlr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) Get(c *gin.Context) {

	// Extraer el parámetro 'id' de la URL
	id := c.Param("id")

	// Consumir servicio
	contentItem, err := h.Service.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if contentItem == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "content item not found"})
		return
	}

	// Traducir respuesta
	c.JSON(http.StatusOK, contentItem)
}
