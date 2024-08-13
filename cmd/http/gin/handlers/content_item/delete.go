package content_item_hdlr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) Delete(c *gin.Context) {

	// Extraer el parámetro 'id' de la URL
	id := c.Param("id")

	// Consumir servicio
	err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Traducir respuesta
	c.Status(http.StatusOK)
}
