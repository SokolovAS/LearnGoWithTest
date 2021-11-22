package main

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var AllRomainNumerals = []RomanNumeral{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range AllRomainNumerals {
		for arabic >= numeral.Value {
			//fmt.Println(numeral.Symbol, numeral.Value, arabic)
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}
