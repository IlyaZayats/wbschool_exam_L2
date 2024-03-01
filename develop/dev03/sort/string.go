package sort

type StringStruct struct {
	Str   string
	Index int
}

type StringSort []StringStruct

func (s StringSort) Len() int {
	return len(s)
}

func (s StringSort) Less(i, j int) bool {
	return s[i].Str < s[j].Str
}

func (s StringSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
