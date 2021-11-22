package main

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Romain string
	}{
		{1, "I"},
		{4, "IV"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Romain), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Romain {
				t.Errorf("got %q, want %q", got, test.Romain)
			}
		})
	}

}
