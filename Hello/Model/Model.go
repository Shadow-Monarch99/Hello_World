package Model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorldResponse struct {
	Output string `json:"Output" example:"Hello World"`
}

// @Summary The Model: access to data
// @Description Accesses database and outputs "Hello World"
// @Tags Model
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
func Out(c *gin.Context) {
	response := HelloWorldResponse{
		Output: "Hello World",
	}
	c.JSON(http.StatusOK, response)
}
