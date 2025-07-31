package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
  if input <= 0 || input > 3999 {
    return "", errors.New("Cannot create a numeral from this number");
  }

  var result string;
  for (input > 0) {
    r, s := calcRoman(input)
    input = r;
    result += s;
  }

  return result, nil;
  
}

func calcRoman(i int) (int, string) {

  if i / 1000 > 0 { return i % 1000, stringRepeat(i / 1000, "M")  }
  if i / 900 > 0 { return i % 900, stringRepeat(i / 900, "CM") }
  if i / 500 > 0 { return i % 500, stringRepeat(i / 500, "D") }
  if i / 400 > 0 { return i % 400, stringRepeat(i / 400, "CD") }
  if i / 100 > 0 { return i % 100, stringRepeat(i / 100, "C") }
  if i / 90 > 0 { return i % 90, stringRepeat(i / 90, "XC") }
  if i / 50 > 0 { return i % 50, stringRepeat(i / 50, "L") }
  if i / 40 > 0 { return i % 40, stringRepeat(i / 40, "XL") }
  if i / 10 > 0 { return i % 10, stringRepeat(i / 10, "X") }
  if i / 9 > 0 { return i % 9, stringRepeat(i / 9, "IX") }
  if i / 5 > 0 { return i % 5, stringRepeat(i / 5, "V") }
  if i / 4 > 0 { return i % 4, stringRepeat(i / 4, "IV") }
  return i % 1, stringRepeat(i / 1, "I")
}

func stringRepeat(i int, s string) string {
  var final string;
  for j := 0; j < i; j++ {
    final += s
  }

  return final
}
