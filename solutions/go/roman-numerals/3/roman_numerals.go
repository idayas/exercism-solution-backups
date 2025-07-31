package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
  if input <= 0 || input > 3999 {
    return "", errors.New("Cannot create a numeral from this number");
  }

  // Go maps do not maintain a specific order when you iterate through them
  // using a for-range loop, and Roman numerals depends on processing in
  // descending order, so we most ensure the order using stricts


  numerals := []struct {
    Value int
    Symbol string
  }{
    {1000, "M"},
    {900, "CM"},
    {500,  "D"},
    {400, "CD"},
    {100,  "C"},
    {90,  "XC"},
    {50,   "L"},
    {40,  "XL"},
    {10,   "X"},
    {9,   "IX"},
    {5,    "V"},
    {4,   "IV"},
    {1,    "I"},
  }


  var result string;
  for _, n := range numerals {
    if input / n.Value > 0 {
      for j := 0; j < input / n.Value; j++ { result += n.Symbol }
      input = input % n.Value
    }
  }

  return result, nil;
  
}

