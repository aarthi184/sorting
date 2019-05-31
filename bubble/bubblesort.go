package bubble

func Sort(array []int) error {
	var (
		tmp     int
		swapped = false
	)
	for i := len(array) - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				tmp = array[j]
				array[j] = array[j-1]
				array[j-1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nil
}
