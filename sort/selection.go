package sort

func Selection(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		minValue := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if minValue > arr[j] {
				minValue = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
