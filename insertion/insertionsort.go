package insertion

func Sort(array []int) (error) {
	var (
		i, j, current int
	)
	for i = 1; i < len(array); i++ {
		current = array[i]
		for j = i; j > 0 && array[j-1] > current; j-- {
			array[j] = array[j-1]
		}
		array[j]= current
	}
	return nil
}
