package controller

import (
	"github.com/MaxFuhrich/gin-calculator/entity"
	"github.com/MaxFuhrich/gin-calculator/service"
	"github.com/gin-gonic/gin"
)

type CalculatorController interface {
	Add(ctx *gin.Context) (*entity.Calc, error)
	Sub(ctx *gin.Context) (*entity.Calc, error)
	Mult(ctx *gin.Context) (*entity.Calc, error)
	Div(ctx *gin.Context) (*entity.Calc, error)
}

type controller struct {
	calculatorService service.CalculatorService
}

func (c controller) Add(ctx *gin.Context) (*entity.Calc, error) {
	var calculation entity.Calc
	err := ctx.BindJSON(&calculation)
	if err != nil {
		return nil, err
	}
	return c.calculatorService.Add(calculation), nil
	//panic("implement me")
}

func (c controller) Sub(ctx *gin.Context) (*entity.Calc, error) {
	var calculation entity.Calc
	err := ctx.BindJSON(&calculation)
	if err != nil {
		return nil, err
	}
	return c.calculatorService.Sub(calculation), nil
}

func (c controller) Mult(ctx *gin.Context) (*entity.Calc, error) {
	var calculation entity.Calc
	err := ctx.BindJSON(&calculation)
	if err != nil {
		return nil, err
	}
	return c.calculatorService.Mult(calculation), nil
}

func (c controller) Div(ctx *gin.Context) (*entity.Calc, error) {
	var calculation entity.Calc
	err := ctx.BindJSON(&calculation)
	if err != nil {
		return nil, err
	}
	return c.calculatorService.Mult(calculation), nil
}

func New(service service.CalculatorService) CalculatorController {
	return &controller{calculatorService: service}

}
