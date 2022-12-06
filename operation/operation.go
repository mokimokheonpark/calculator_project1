package operation

import (
	"errors"
	"math"
)

var errOverflow = errors.New("float64 overflow")

// Number: struct and method

type Number struct {
	Value float64
}

func (number Number) Operate() (float64, error) {
	return number.Value, nil
}

// Addition: struct and method

type Addition struct {
	LeftNumber  float64
	RightNumber float64
}

func (addition Addition) Operate() (float64, error) {
	if (addition.LeftNumber > 0 && addition.RightNumber > math.MaxFloat64-addition.LeftNumber) ||
		(addition.LeftNumber < 0 && addition.RightNumber < -math.MaxFloat64-addition.LeftNumber) {
		return 0, errOverflow
	}
	return addition.LeftNumber + addition.RightNumber, nil
}

// Subtraction: struct and method

type Subtraction struct {
	LeftNumber  float64
	RightNumber float64
}

func (subtraction Subtraction) Operate() (float64, error) {
	if (subtraction.LeftNumber > 0 && subtraction.RightNumber < subtraction.LeftNumber-math.MaxFloat64) ||
		(subtraction.LeftNumber < 0 && subtraction.RightNumber > subtraction.LeftNumber+math.MaxFloat64) {
		return 0, errOverflow
	}
	return subtraction.LeftNumber - subtraction.RightNumber, nil
}

// Multiplication: struct and method

type Multiplication struct {
	LeftNumber  float64
	RightNumber float64
}

func (multiplication Multiplication) Operate() (float64, error) {
	if multiplication.LeftNumber == 0 || multiplication.RightNumber == 0 {
		return 0, nil
	}
	if math.Abs(multiplication.LeftNumber) > math.Abs(math.MaxFloat64/multiplication.RightNumber) {
		return 0, errOverflow
	}
	return multiplication.LeftNumber * multiplication.RightNumber, nil
}

// Division: struct and method

type Division struct {
	LeftNumber  float64
	RightNumber float64
}

func (division Division) Operate() (float64, error) {
	if division.RightNumber == 0 {
		return 0, errors.New("division by zero")
	}
	if division.LeftNumber == 0 {
		return 0, nil
	}
	return division.LeftNumber / division.RightNumber, nil
}

// Expression: interface of methods

type Expression interface {
	Operate() (float64, error)
}

// GetExpression: given a valid operator and two numbers, get the corresponding Expression

func GetExpression(operator string, leftNum float64, rightNum float64) (Expression, error) {
	switch operator {
	case "+":
		return Addition{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "-":
		return Subtraction{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "*":
		return Multiplication{LeftNumber: leftNum, RightNumber: rightNum}, nil
	case "/":
		return Division{LeftNumber: leftNum, RightNumber: rightNum}, nil
	default:
		return nil, errors.New("invalid operator")
	}
}
