package service

import (
	"fmt"
	"github.com/MaxFuhrich/gin-calculator/entity"
)

type CalculatorService interface {
	Add(calc entity.Calc) *entity.Calc
	Sub(calc entity.Calc) *entity.Calc
	Mult(calc entity.Calc) *entity.Calc
	Div(calc entity.Calc) *entity.Calc
}

type calculatorService struct {
}

func (c *calculatorService) Add(calc entity.Calc) *entity.Calc {
	calc.Result = calc.A + calc.B
	calc.OperationUsed = fmt.Sprintf("%f + %f = %f", calc.A, calc.B, calc.Result)
	return &calc
}

func (c *calculatorService) Sub(calc entity.Calc) *entity.Calc {
	calc.Result = calc.A - calc.B
	calc.OperationUsed = fmt.Sprintf("%f - %f = %f", calc.A, calc.B, calc.Result)
	return &calc
}

func (c *calculatorService) Mult(calc entity.Calc) *entity.Calc {
	calc.Result = calc.A * calc.B
	calc.OperationUsed = fmt.Sprintf("%f * %f = %f", calc.A, calc.B, calc.Result)
	return &calc
}

func (c *calculatorService) Div(calc entity.Calc) *entity.Calc {
	calc.Result = calc.A / calc.B
	calc.OperationUsed = fmt.Sprintf("%f / %f = %f", calc.A, calc.B, calc.Result)
	return &calc
}

func New() CalculatorService {
	return &calculatorService{}
}
