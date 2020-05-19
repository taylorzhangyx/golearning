package magic

var magics []int = []int{2, 3, 5, 7, 11}

// Magic returns if has the matic number
func Magic(nums ...int) bool {
	for _, n := range nums {
		for _, m := range magics {
			if m == n {
				return true
			}
		}
	}
	return false
}
