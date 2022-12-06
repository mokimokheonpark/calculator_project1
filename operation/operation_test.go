package operation_test

import (
	"calculator_project/operation"
	"testing"
)

var err string = "Wrong Result"

// Unit Test of Number.Operate Function

func TestNumberOperate(t *testing.T) {
	number1, _ := operation.Number{Value: 5}.Operate()
	if number1 != 5 {
		t.Error(err)
	}

	number2, _ := operation.Number{Value: -12.345}.Operate()
	if number2 != -12.345 {
		t.Error(err)
	}

	number3, _ := operation.Number{Value: 0}.Operate()
	if number3 != 0 {
		t.Error(err)
	}
}

// Unit Test of Addition.Operate Function

func TestAdditionOperate(t *testing.T) {
	addition1, _ := operation.Addition{LeftNumber: 1, RightNumber: 2}.Operate()
	if addition1 != 3 {
		t.Error(err)
	}

	addition2, _ := operation.Addition{LeftNumber: 6, RightNumber: -1}.Operate()
	if addition2 != 5 {
		t.Error(err)
	}

	addition3, _ := operation.Addition{LeftNumber: -100, RightNumber: 15}.Operate()
	if addition3 != -85 {
		t.Error(err)
	}

	addition4, _ := operation.Addition{LeftNumber: -1, RightNumber: -10000}.Operate()
	if addition4 != -10001 {
		t.Error(err)
	}

	addition5, _ := operation.Addition{LeftNumber: 0, RightNumber: 0}.Operate()
	if addition5 != 0 {
		t.Error(err)
	}

	addition6, _ := operation.Addition{LeftNumber: 7, RightNumber: 0}.Operate()
	if addition6 != 7 {
		t.Error(err)
	}

	addition7, _ := operation.Addition{LeftNumber: 0, RightNumber: -2}.Operate()
	if addition7 != -2 {
		t.Error(err)
	}

	addition8, _ := operation.Addition{LeftNumber: -5, RightNumber: 5}.Operate()
	if addition8 != 0 {
		t.Error(err)
	}

	addition9, _ := operation.Addition{LeftNumber: 1, RightNumber: 0.1}.Operate()
	if addition9 != 1.1 {
		t.Error(err)
	}

	addition10, _ := operation.Addition{LeftNumber: -0.001, RightNumber: 0}.Operate()
	if addition10 != -0.001 {
		t.Error(err)
	}

	addition11, _ := operation.Addition{LeftNumber: 1.234, RightNumber: -5.678}.Operate()
	if addition11 != -4.444 {
		t.Error(err)
	}

	addition12, _ := operation.Addition{LeftNumber: -2.5, RightNumber: -7.5}.Operate()
	if addition12 != -10 {
		t.Error(err)
	}

	addition13, _ := operation.Addition{LeftNumber: 5.00, RightNumber: 6.00000}.Operate()
	if addition13 != 11 {
		t.Error(err)
	}

	_, addition14Error := operation.Addition{LeftNumber: 1.7e+308, RightNumber: 1.2e+308}.Operate()
	if addition14Error == nil {
		t.Error(err)
	}

	_, addition15Error := operation.Addition{LeftNumber: -1.5e+308, RightNumber: -0.3e+308}.Operate()
	if addition15Error == nil {
		t.Error(err)
	}
}

// Unit Test of Subtraction.Operate Function

func TestSubtractionOperate(t *testing.T) {
	subtraction1, _ := operation.Subtraction{LeftNumber: 7, RightNumber: 2}.Operate()
	if subtraction1 != 5 {
		t.Error(err)
	}

	subtraction2, _ := operation.Subtraction{LeftNumber: 36, RightNumber: -3}.Operate()
	if subtraction2 != 39 {
		t.Error(err)
	}

	subtraction3, _ := operation.Subtraction{LeftNumber: -100, RightNumber: 11}.Operate()
	if subtraction3 != -111 {
		t.Error(err)
	}

	subtraction4, _ := operation.Subtraction{LeftNumber: -10000, RightNumber: -8800}.Operate()
	if subtraction4 != -1200 {
		t.Error(err)
	}

	subtraction5, _ := operation.Subtraction{LeftNumber: 0, RightNumber: 0}.Operate()
	if subtraction5 != 0 {
		t.Error(err)
	}

	subtraction6, _ := operation.Subtraction{LeftNumber: 6, RightNumber: 0}.Operate()
	if subtraction6 != 6 {
		t.Error(err)
	}

	subtraction7, _ := operation.Subtraction{LeftNumber: 0, RightNumber: -22}.Operate()
	if subtraction7 != 22 {
		t.Error(err)
	}

	subtraction8, _ := operation.Subtraction{LeftNumber: 5, RightNumber: 5}.Operate()
	if subtraction8 != 0 {
		t.Error(err)
	}

	subtraction9, _ := operation.Subtraction{LeftNumber: 2, RightNumber: 0.3}.Operate()
	if subtraction9 != 1.7 {
		t.Error(err)
	}

	subtraction10, _ := operation.Subtraction{LeftNumber: -0.0000009, RightNumber: 0}.Operate()
	if subtraction10 != -0.0000009 {
		t.Error(err)
	}

	subtraction11, _ := operation.Subtraction{LeftNumber: 1.2345, RightNumber: -6.789}.Operate()
	if subtraction11 != 8.0235 {
		t.Error(err)
	}

	subtraction12, _ := operation.Subtraction{LeftNumber: -3.5, RightNumber: -6.5}.Operate()
	if subtraction12 != 3 {
		t.Error(err)
	}

	subtraction13, _ := operation.Subtraction{LeftNumber: 10.00, RightNumber: 4.0000000000}.Operate()
	if subtraction13 != 6 {
		t.Error(err)
	}

	_, subtraction14Error := operation.Subtraction{LeftNumber: 1.6e+308, RightNumber: -1.5e+308}.Operate()
	if subtraction14Error == nil {
		t.Error(err)
	}

	_, subtraction15Error := operation.Subtraction{LeftNumber: -1.7e+308, RightNumber: 0.1e+308}.Operate()
	if subtraction15Error == nil {
		t.Error(err)
	}
}

// Unit Test of Multiplication.Operate Function

func TestMultiplicationOperate(t *testing.T) {
	multiplication1, _ := operation.Multiplication{LeftNumber: 2, RightNumber: 4}.Operate()
	if multiplication1 != 8 {
		t.Error(err)
	}

	multiplication2, _ := operation.Multiplication{LeftNumber: 5, RightNumber: -3}.Operate()
	if multiplication2 != -15 {
		t.Error(err)
	}

	multiplication3, _ := operation.Multiplication{LeftNumber: -100, RightNumber: 12}.Operate()
	if multiplication3 != -1200 {
		t.Error(err)
	}

	multiplication4, _ := operation.Multiplication{LeftNumber: -20, RightNumber: -4000}.Operate()
	if multiplication4 != 80000 {
		t.Error(err)
	}

	multiplication5, _ := operation.Multiplication{LeftNumber: 0, RightNumber: 0}.Operate()
	if multiplication5 != 0 {
		t.Error(err)
	}

	multiplication6, _ := operation.Multiplication{LeftNumber: 9, RightNumber: 0}.Operate()
	if multiplication6 != 0 {
		t.Error(err)
	}

	multiplication7, _ := operation.Multiplication{LeftNumber: 0, RightNumber: -123456789}.Operate()
	if multiplication7 != 0 {
		t.Error(err)
	}

	multiplication8, _ := operation.Multiplication{LeftNumber: 100, RightNumber: 100}.Operate()
	if multiplication8 != 10000 {
		t.Error(err)
	}

	multiplication9, _ := operation.Multiplication{LeftNumber: 99999999, RightNumber: 1}.Operate()
	if multiplication9 != 99999999 {
		t.Error(err)
	}

	multiplication10, _ := operation.Multiplication{LeftNumber: 10, RightNumber: 0.3}.Operate()
	if multiplication10 != 3 {
		t.Error(err)
	}

	multiplication11, _ := operation.Multiplication{LeftNumber: -0.001, RightNumber: 400000}.Operate()
	if multiplication11 != -400 {
		t.Error(err)
	}

	multiplication12, _ := operation.Multiplication{LeftNumber: 3.5, RightNumber: -4.5}.Operate()
	if multiplication12 != -15.75 {
		t.Error(err)
	}

	multiplication13, _ := operation.Multiplication{LeftNumber: -10.1, RightNumber: -0.111}.Operate()
	if multiplication13 != 1.1211 {
		t.Error(err)
	}

	multiplication14, _ := operation.Multiplication{LeftNumber: 9.00, RightNumber: 8.00000}.Operate()
	if multiplication14 != 72 {
		t.Error(err)
	}

	_, multiplication15Error := operation.Multiplication{LeftNumber: 1.79e+308, RightNumber: 100}.Operate()
	if multiplication15Error == nil {
		t.Error(err)
	}

	_, multiplication16Error := operation.Multiplication{LeftNumber: 1.8e+156, RightNumber: -1.8e+156}.Operate()
	if multiplication16Error == nil {
		t.Error(err)
	}

	_, multiplication17Error := operation.Multiplication{LeftNumber: -1.0e+200, RightNumber: 1.2e+109}.Operate()
	if multiplication17Error == nil {
		t.Error(err)
	}

	_, multiplication18Error := operation.Multiplication{LeftNumber: -12.34, RightNumber: -1.797e+308}.Operate()
	if multiplication18Error == nil {
		t.Error(err)
	}
}

// Unit Test of Division.Operate Function

func TestDivisionOperate(t *testing.T) {
	division1, _ := operation.Division{LeftNumber: 6, RightNumber: 2}.Operate()
	if division1 != 3 {
		t.Error(err)
	}

	division2, _ := operation.Division{LeftNumber: 15, RightNumber: -3}.Operate()
	if division2 != -5 {
		t.Error(err)
	}

	division3, _ := operation.Division{LeftNumber: -100, RightNumber: 8}.Operate()
	if division3 != -12.5 {
		t.Error(err)
	}

	division4, _ := operation.Division{LeftNumber: -3, RightNumber: -60000}.Operate()
	if division4 != 0.00005 {
		t.Error(err)
	}

	division5, _ := operation.Division{LeftNumber: 0, RightNumber: 123456789}.Operate()
	if division5 != 0 {
		t.Error(err)
	}

	division6, _ := operation.Division{LeftNumber: 987654321, RightNumber: 1}.Operate()
	if division6 != 987654321 {
		t.Error(err)
	}

	division7, _ := operation.Division{LeftNumber: 7777777, RightNumber: 7777777}.Operate()
	if division7 != 1 {
		t.Error(err)
	}

	division8, _ := operation.Division{LeftNumber: 90, RightNumber: 0.5}.Operate()
	if division8 != 180 {
		t.Error(err)
	}

	division9, _ := operation.Division{LeftNumber: -0.03, RightNumber: 200}.Operate()
	if division9 != -0.00015 {
		t.Error(err)
	}

	division10, _ := operation.Division{LeftNumber: 10.5, RightNumber: -3.5}.Operate()
	if division10 != -3 {
		t.Error(err)
	}

	division11, _ := operation.Division{LeftNumber: -99.99, RightNumber: -6.6}.Operate()
	if division11 != 15.15 {
		t.Error(err)
	}

	division12, _ := operation.Division{LeftNumber: 125.00, RightNumber: 5.00000}.Operate()
	if division12 != 25 {
		t.Error(err)
	}

	_, division13Error := operation.Division{LeftNumber: 21, RightNumber: 0}.Operate()
	if division13Error == nil {
		t.Error(err)
	}

	_, division14Error := operation.Division{LeftNumber: 0, RightNumber: -0}.Operate()
	if division14Error == nil {
		t.Error(err)
	}
}

// Unit Test of GetExpression Function

func TestGetExpression(t *testing.T) {
	getExpression1, _ := operation.GetExpression("+", 1, 2)
	var expression1 operation.Expression
	expression1 = operation.Addition{LeftNumber: 1, RightNumber: 2}
	if getExpression1 != expression1 {
		t.Error(err)
	}

	getExpression2, _ := operation.GetExpression("-", 9, 7)
	var expression2 operation.Expression
	expression2 = operation.Subtraction{LeftNumber: 9, RightNumber: 7}
	if getExpression2 != expression2 {
		t.Error(err)
	}

	getExpression3, _ := operation.GetExpression("*", 3, 5)
	var expression3 operation.Expression
	expression3 = operation.Multiplication{LeftNumber: 3, RightNumber: 5}
	if getExpression3 != expression3 {
		t.Error(err)
	}

	getExpression4, _ := operation.GetExpression("/", 8, 4)
	var expression4 operation.Expression
	expression4 = operation.Division{LeftNumber: 8, RightNumber: 4}
	if getExpression4 != expression4 {
		t.Error(err)
	}

	_, getExpression5Error := operation.GetExpression("^", 2, 10)
	if getExpression5Error == nil {
		t.Error(err)
	}

	_, getExpression6Error := operation.GetExpression("%", 20, 6)
	if getExpression6Error == nil {
		t.Error(err)
	}
}
