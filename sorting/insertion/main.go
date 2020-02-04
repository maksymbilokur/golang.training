package main
func insertionSort(array []int) {
	var n = len(array)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}

}