package sort

type NumericStruct struct {
	Num   float64
	Index int
}

type NumericSort []NumericStruct

func (n NumericSort) Len() int {
	return len(n)
}

func (n NumericSort) Less(i, j int) bool {
	return n[i].Num < n[j].Num
}

func (n NumericSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
