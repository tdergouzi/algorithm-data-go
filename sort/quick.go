package sort

func Quick(arr []int64) []int64 {
	// p pivot r
	quick(arr, 0, len(arr)-1)
	return arr
}

func quick(arr []int64, p, q int) {
	if q-p < 1 {
		return
	}
	r := partition(arr, p, q)
	quick(arr, p, r-1)
	quick(arr, r+1, q)
}

func partition(arr []int64, p, q int) int {
	i := p
	pivot := arr[q]
	for j := p; j < q; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[q] = arr[q], arr[i]
	return i
}
