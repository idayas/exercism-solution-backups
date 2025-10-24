package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func SeeIfTriangle(a, b, c float64) bool {
	if a + b >= c && b + c >= a && a + c >= b && a + b + c > 0 {
		return true
	}
	return false
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !SeeIfTriangle(a, b, c) {
		return NaT
	}
	
	equalSides := 0;
	if a == b { equalSides++ }
	if b == c { equalSides++ }
	if c == a { equalSides++ }

	switch equalSides {
	case 3:
		k = Equ
	case 1:
		k = Iso
	case 0:
		k = Sca
	default:
		k = NaT
	}

	return k
}
