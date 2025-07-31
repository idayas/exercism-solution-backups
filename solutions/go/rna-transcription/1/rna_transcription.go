package strand

func ToRNA(dna string) string {
  mapping := map[rune]rune{
    'G': 'C',
    'C': 'G',
    'T': 'A',
    'A': 'U',
  }
  rna := []rune{}
  for _, c := range dna {
    rna = append(rna, mapping[c])
  }
  return string(rna)
}
