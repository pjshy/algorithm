package algorithm

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, value := range nums {
		l := target - value
		r, ok := m[l]
		if ok {
			return []int{r, k}
		}
		m[l] = k
	}
	return []int{}
}
