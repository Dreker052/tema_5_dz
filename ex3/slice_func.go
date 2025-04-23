package ex3

func SliceFunc(sl []int) int {
	var sum int
	for _, v := range sl {
		sum += v
	}
	return sum
}
