package main

import (
	"fmt"
	"github.com/MaxFuhrich/gin-calculator/controller"
	"github.com/MaxFuhrich/gin-calculator/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//import "github.com/gin-gonic/gin"

func main() {
	calculatorService := service.New()
	calculatorController := controller.New(calculatorService)
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.POST("/add", func(context *gin.Context) {
		calculation, err := calculatorController.Add(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, calculation)
		}
	})
	server.POST("/sub", func(context *gin.Context) {
		calculation, err := calculatorController.Sub(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, calculation)
		}
	})
	server.POST("/mult", func(context *gin.Context) {
		calculation, err := calculatorController.Mult(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, calculation)
		}
	})
	server.POST("/div", func(context *gin.Context) {
		calculation, err := calculatorController.Div(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, calculation)
		}
	})
	fmt.Println(server.Run(":8080"))
	/*var calc service.CalculatorService
	bla := calc.Add()*/
}
