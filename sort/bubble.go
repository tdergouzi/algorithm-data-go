package sort

func Bubble(arr []int64) []int64 {
	bubble(arr)
	return arr
}

func bubble(arr []int64) {
	for i := 0; i < len(arr)-1; i++ {
		existSwap := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				existSwap = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !existSwap {
			break // 单次冒泡，没有发生交换事件，数组为满有序度数组
		}
	}
}
