package sort

// 归并排序
func Merge(arr []int64) []int64 {
	MergeS(arr, 0, len(arr))
	return arr
}

func MergeS(arr []int64, s, e int) {
	// 递归终止判断
	if e-s <= 1 {
		return
	}
	m := (s + e) / 2
	// 递归向下
	MergeS(arr, s, m)
	MergeS(arr, m, e)
	// 归并向上
	MergeM(arr, s, m, e)
}

func MergeM(arr []int64, s, m, e int) {
	k := s
	i := 0
	j := 0
	arr1 := make([]int64, m-s)
	copy(arr1, arr[s:m])
	arr2 := make([]int64, e-m)
	copy(arr2, arr[m:e])

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}

	if i >= len(arr1) {
		for l := j; l < len(arr2); l++ {
			arr[k] = arr2[l]
			k++
		}
	} else {
		for l := i; l < len(arr1); l++ {
			arr[k] = arr1[l]
			k++
		}
	}
}
