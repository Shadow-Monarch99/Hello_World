package Controller

import (
	"github.com/gin-gonic/gin"
	"main.go/Model"
)

// out godoc
// @Summary Handler for localhost:8080/HelloWorld
// @Description Handles http request,Asks data from model and returns "Hello World"
// @Tags Controller
// @Accept json
// @Produce json
// @Success 200 {object} Model.HelloWorldResponse
// @Router /HelloWorld [get]
func HelloWorld(c *gin.Context) {
	Model.Out(c)
}
