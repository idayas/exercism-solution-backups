package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
  counter := 0
  if n < 1 {
    return 0, errors.New("Number must be greater than one")
  }

  for n > 1 {
    if n % 2 == 0 {
      n = n / 2
    } else {
      n = (3*n) + 1 
    }
    counter++
  }
  return counter, nil
}
