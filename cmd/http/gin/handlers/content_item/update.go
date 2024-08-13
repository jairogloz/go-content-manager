package content_item_hdlr

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (h *HttpHandler) Update(c *gin.Context) {

	// Traducir request
	var contentItemUpdateParams domain.ContentItemUpdateParams
	if err := c.ShouldBindJSON(&contentItemUpdateParams); err != nil {
		log.Println("Failed to bind JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extraer el parámetro 'id' de la URL
	id := c.Param("id")

	// Consumir servicio
	updateCount, err := h.Service.Update(id, contentItemUpdateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if updateCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "content item not found"})
		return
	}

	// Traducir respuesta
	c.JSON(http.StatusOK, gin.H{"updated": updateCount})
}
