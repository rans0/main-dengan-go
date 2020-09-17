package sort

type Miner interface {
	Len() int
	Less(i, j int) bool
	ElemIx(ix int) interface{}
}

func Min (data Miner) interface{} {
	min := data.ElemIx(0)
	for i:=1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemIx(i)
		}
	}
	return min
}

type IntArray []int
func (p IntArray) Len() int{
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) ElemIx(ix int) interface{} {
	return p[ix]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringArray) ElemIx(ix int) interface{} {
	return p[ix]
}