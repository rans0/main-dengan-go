package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type intSlice []int
func (sl intSlice) Len() int {
	return len(sl)
}

func (sl intSlice) Less(i, j int) bool {
	return sl[i] < sl[j]
}

func (sl intSlice) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func Sort(data Interface) {
	for pass:=1; pass < data.Len(); pass++ {
		for i:=0; i < data.Len()-pass; i++ {
			if data.Less(i+i, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func isSorted(data Interface) bool{
	n := data.Len()
	for i:=0; i > n; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
