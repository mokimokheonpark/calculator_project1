package parser_test

import (
	"calculator_project/operation"
	"calculator_project/parser"
	"reflect"
	"testing"
)

var err string = "Wrong Result"

// Unit Test of GetTokens Function

func TestGetTokens(t *testing.T) {
	tokens1 := parser.GetTokens("")
	if len(tokens1) != 0 {
		t.Error(err)
	}

	tokens2 := parser.GetTokens("123456")
	if !reflect.DeepEqual(tokens2, []string{"123456"}) {
		t.Error(err)
	}

	tokens3 := parser.GetTokens("   -7.89")
	if !reflect.DeepEqual(tokens3, []string{"-7.89"}) {
		t.Error(err)
	}

	tokens4 := parser.GetTokens("+     ")
	if !reflect.DeepEqual(tokens4, []string{"+"}) {
		t.Error(err)
	}

	tokens5 := parser.GetTokens("Ab3?")
	if !reflect.DeepEqual(tokens5, []string{"Ab3?"}) {
		t.Error(err)
	}

	tokens6 := parser.GetTokens("1+")
	if !reflect.DeepEqual(tokens6, []string{"1", "+"}) {
		t.Error(err)
	}

	tokens7 := parser.GetTokens("2 -")
	if !reflect.DeepEqual(tokens7, []string{"2", "-"}) {
		t.Error(err)
	}

	tokens8 := parser.GetTokens(" -3  *     ")
	if !reflect.DeepEqual(tokens8, []string{"-3", "*"}) {
		t.Error(err)
	}

	tokens9 := parser.GetTokens("4.5/")
	if !reflect.DeepEqual(tokens9, []string{"4.5", "/"}) {
		t.Error(err)
	}

	tokens10 := parser.GetTokens("123 4567")
	if !reflect.DeepEqual(tokens10, []string{"123", "4567"}) {
		t.Error(err)
	}

	tokens11 := parser.GetTokens("+ 4567")
	if !reflect.DeepEqual(tokens11, []string{"+", "4567"}) {
		t.Error(err)
	}

	tokens12 := parser.GetTokens("+-")
	if !reflect.DeepEqual(tokens12, []string{"+", "-"}) {
		t.Error(err)
	}

	tokens13 := parser.GetTokens("ABC   abcde")
	if !reflect.DeepEqual(tokens13, []string{"ABC", "abcde"}) {
		t.Error(err)
	}

	tokens14 := parser.GetTokens("HelloWorld*")
	if !reflect.DeepEqual(tokens14, []string{"HelloWorld", "*"}) {
		t.Error(err)
	}

	tokens15 := parser.GetTokens("1 + 2")
	if !reflect.DeepEqual(tokens15, []string{"1", "+", "2"}) {
		t.Error(err)
	}

	tokens16 := parser.GetTokens("2--5")
	if !reflect.DeepEqual(tokens16, []string{"2", "-", "-5"}) {
		t.Error(err)
	}

	tokens17 := parser.GetTokens("   -3     *       -7  ")
	if !reflect.DeepEqual(tokens17, []string{"-3", "*", "-7"}) {
		t.Error(err)
	}

	tokens18 := parser.GetTokens("+12.48 / +2.55")
	if !reflect.DeepEqual(tokens18, []string{"+12.48", "/", "+2.55"}) {
		t.Error(err)
	}

	tokens19 := parser.GetTokens(" ABC+123   ")
	if !reflect.DeepEqual(tokens19, []string{"ABC", "+", "123"}) {
		t.Error(err)
	}

	tokens20 := parser.GetTokens("812376   0.11  0")
	if !reflect.DeepEqual(tokens20, []string{"812376", "0.11", "0"}) {
		t.Error(err)
	}

	tokens21 := parser.GetTokens("7 +B")
	if !reflect.DeepEqual(tokens21, []string{"7", "+", "B"}) {
		t.Error(err)
	}

	tokens22 := parser.GetTokens("+ Z 1")
	if !reflect.DeepEqual(tokens22, []string{"+", "Z", "1"}) {
		t.Error(err)
	}

	tokens23 := parser.GetTokens("x - y")
	if !reflect.DeepEqual(tokens23, []string{"x", "-", "y"}) {
		t.Error(err)
	}

	tokens24 := parser.GetTokens("+12.1212 token *")
	if !reflect.DeepEqual(tokens24, []string{"+12.1212", "token", "*"}) {
		t.Error(err)
	}

	tokens25 := parser.GetTokens("1 + 2 -3")
	if !reflect.DeepEqual(tokens25, []string{"1", "+", "2", "-3"}) {
		t.Error(err)
	}

	tokens26 := parser.GetTokens("2** 10")
	if !reflect.DeepEqual(tokens26, []string{"2", "*", "*", "10"}) {
		t.Error(err)
	}

	tokens27 := parser.GetTokens("0.5 / 2.3 *")
	if !reflect.DeepEqual(tokens27, []string{"0.5", "/", "2.3", "*"}) {
		t.Error(err)
	}

	tokens28 := parser.GetTokens("--2--- -")
	if !reflect.DeepEqual(tokens28, []string{"-", "-", "2---", "-"}) {
		t.Error(err)
	}

	tokens29 := parser.GetTokens("10 + 2 + 3.7")
	if !reflect.DeepEqual(tokens29, []string{"10", "+", "2", "+", "3.7"}) {
		t.Error(err)
	}

	tokens30 := parser.GetTokens("   a + 10-5.7 * % A3b  / gggg  ")
	if !reflect.DeepEqual(tokens30, []string{"a", "+", "10-5.7", "*", "%", "A3b", "/", "gggg"}) {
		t.Error(err)
	}
}

// Unit Test of GetExpressionWithOneToken Function

func TestGetExpressionWithOneToken(t *testing.T) {
	getExpressionWithOneToken1, _ := parser.GetExpressionWithOneToken("7")
	var expressionWithOneToken1 operation.Expression = operation.Number{Value: 7}
	if getExpressionWithOneToken1 != expressionWithOneToken1 {
		t.Error(err)
	}

	getExpressionWithOneToken2, _ := parser.GetExpressionWithOneToken("-3.45")
	var expressionWithOneToken2 operation.Expression = operation.Number{Value: -3.45}
	if getExpressionWithOneToken2 != expressionWithOneToken2 {
		t.Error(err)
	}

	_, getExpressionWithOneToken3Error := parser.GetExpressionWithOneToken("A")
	if getExpressionWithOneToken3Error == nil {
		t.Error(err)
	}

	_, getExpressionWithOneToken4Error := parser.GetExpressionWithOneToken("+")
	if getExpressionWithOneToken4Error == nil {
		t.Error(err)
	}

	_, getExpressionWithOneToken5Error := parser.GetExpressionWithOneToken("-123abc")
	if getExpressionWithOneToken5Error == nil {
		t.Error(err)
	}

	_, getExpressionWithOneToken6Error := parser.GetExpressionWithOneToken("0.987d")
	if getExpressionWithOneToken6Error == nil {
		t.Error(err)
	}
}

// Unit Test of GetExpressionWithThreeTokens Function

func TestGetExpressionWithThreeTokens(t *testing.T) {
	getExpressionWithThreeTokens1, _ := parser.GetExpressionWithThreeTokens("1", "+", "2")
	var expressionWithThreeTokens1 operation.Expression = operation.Addition{LeftNumber: 1, RightNumber: 2}
	if getExpressionWithThreeTokens1 != expressionWithThreeTokens1 {
		t.Error(err)
	}

	getExpressionWithThreeTokens2, _ := parser.GetExpressionWithThreeTokens("200", "-", "13.5")
	var expressionWithThreeTokens2 operation.Expression = operation.Subtraction{LeftNumber: 200, RightNumber: 13.5}
	if getExpressionWithThreeTokens2 != expressionWithThreeTokens2 {
		t.Error(err)
	}

	getExpressionWithThreeTokens3, _ := parser.GetExpressionWithThreeTokens("123.45", "*", "-1000000")
	var expressionWithThreeTokens3 operation.Expression = operation.Multiplication{LeftNumber: 123.45, RightNumber: -1000000}
	if getExpressionWithThreeTokens3 != expressionWithThreeTokens3 {
		t.Error(err)
	}

	getExpressionWithThreeTokens4, _ := parser.GetExpressionWithThreeTokens("-0.01", "/", "-99.9999999")
	var expressionWithThreeTokens4 operation.Expression = operation.Division{LeftNumber: -0.01, RightNumber: -99.9999999}
	if getExpressionWithThreeTokens4 != expressionWithThreeTokens4 {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens5Error := parser.GetExpressionWithThreeTokens("a", "*", "10")
	if getExpressionWithThreeTokens5Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens6Error := parser.GetExpressionWithThreeTokens("3232.3232", "%", "-11.11")
	if getExpressionWithThreeTokens6Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens7Error := parser.GetExpressionWithThreeTokens("-100", "+", "ABCDE")
	if getExpressionWithThreeTokens7Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens8Error := parser.GetExpressionWithThreeTokens("1", "2", "3")
	if getExpressionWithThreeTokens8Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens9Error := parser.GetExpressionWithThreeTokens("x", "y", "0")
	if getExpressionWithThreeTokens9Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens10Error := parser.GetExpressionWithThreeTokens("A", "-", "B")
	if getExpressionWithThreeTokens10Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens11Error := parser.GetExpressionWithThreeTokens("10", "hello", "/")
	if getExpressionWithThreeTokens11Error == nil {
		t.Error(err)
	}

	_, getExpressionWithThreeTokens12Error := parser.GetExpressionWithThreeTokens("+", "80", "*")
	if getExpressionWithThreeTokens12Error == nil {
		t.Error(err)
	}
}

// Unit Test of Parser Function

func TestParser(t *testing.T) {
	parser1, _ := parser.Parser([]string{"123456"})
	var expression1 operation.Expression = operation.Number{Value: 123456}
	if parser1 != expression1 {
		t.Error(err)
	}

	parser2, _ := parser.Parser([]string{"-7.89"})
	var expression2 operation.Expression = operation.Number{Value: -7.89}
	if parser2 != expression2 {
		t.Error(err)
	}

	parser3, _ := parser.Parser([]string{"1", "+", "2"})
	var expression3 operation.Expression = operation.Addition{LeftNumber: 1, RightNumber: 2}
	if parser3 != expression3 {
		t.Error(err)
	}

	parser4, _ := parser.Parser([]string{"2", "-", "-5"})
	var expression4 operation.Expression = operation.Subtraction{LeftNumber: 2, RightNumber: -5}
	if parser4 != expression4 {
		t.Error(err)
	}

	parser5, _ := parser.Parser([]string{"-3", "*", "-7"})
	var expression5 operation.Expression = operation.Multiplication{LeftNumber: -3, RightNumber: -7}
	if parser5 != expression5 {
		t.Error(err)
	}

	parser6, _ := parser.Parser([]string{"+12.48", "/", "+2.55"})
	var expression6 operation.Expression = operation.Division{LeftNumber: 12.48, RightNumber: 2.55}
	if parser6 != expression6 {
		t.Error(err)
	}

	_, parser7Error := parser.Parser([]string{})
	if parser7Error == nil {
		t.Error(err)
	}

	_, parser8Error := parser.Parser([]string{"+"})
	if parser8Error == nil {
		t.Error(err)
	}

	_, parser9Error := parser.Parser([]string{"Ab3?"})
	if parser9Error == nil {
		t.Error(err)
	}

	_, parser10Error := parser.Parser([]string{"2", "-"})
	if parser10Error == nil {
		t.Error(err)
	}

	_, parser11Error := parser.Parser([]string{"123", "4567"})
	if parser11Error == nil {
		t.Error(err)
	}

	_, parser12Error := parser.Parser([]string{"+", "4567"})
	if parser12Error == nil {
		t.Error(err)
	}

	_, parser13Error := parser.Parser([]string{"ABC", "+", "123"})
	if parser13Error == nil {
		t.Error(err)
	}

	_, parser14Error := parser.Parser([]string{"812376", "0.11", "0"})
	if parser14Error == nil {
		t.Error(err)
	}

	_, parser15Error := parser.Parser([]string{"7", "+", "B"})
	if parser15Error == nil {
		t.Error(err)
	}

	_, parser16Error := parser.Parser([]string{"+", "Z", "1"})
	if parser16Error == nil {
		t.Error(err)
	}

	_, parser17Error := parser.Parser([]string{"x", "-", "y"})
	if parser17Error == nil {
		t.Error(err)
	}

	_, parser18Error := parser.Parser([]string{"+12.1212", "token", "*"})
	if parser18Error == nil {
		t.Error(err)
	}

	_, parser19Error := parser.Parser([]string{"*", "X", "/"})
	if parser19Error == nil {
		t.Error(err)
	}

	_, parser20Error := parser.Parser([]string{"2", "*", "*", "10"})
	if parser20Error == nil {
		t.Error(err)
	}

	_, parser21Error := parser.Parser([]string{"10", "+", "2", "+", "3.7"})
	if parser21Error == nil {
		t.Error(err)
	}
}
