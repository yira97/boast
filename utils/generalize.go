package utils

func SliI64FromSliI(sliI []int) (sliI64 []int64) {
	sliI64 = make([]int64, len(sliI))
	for _, i := range sliI {
		sliI64 = append(sliI64, int64(i))
	}
	return
}