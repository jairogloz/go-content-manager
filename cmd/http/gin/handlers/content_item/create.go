package content_item_hdlr

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-content-manager/pkg/domain"
)

func (h *HttpHandler) Create(c *gin.Context) {
	// Traducir request
	var contentItemCreateParams domain.ContentItemCreateParams
	if err := c.ShouldBindJSON(&contentItemCreateParams); err != nil {
		log.Println("Failed to bind JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asumimos que el user.ID es "1"
	userID := "1"

	// Consumir servicio
	contentItem, err := h.Service.Create(userID, contentItemCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Traducir respuesta
	c.JSON(200, contentItem)
}
