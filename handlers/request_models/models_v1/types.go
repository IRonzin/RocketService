package models_v1

type CalculatorAddParams struct {
	X float64 `uri:"x" binding:"required"`
	Y float64 `uri:"y" binding:"required"`
}
