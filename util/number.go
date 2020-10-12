package util

import (
	"math/rand"
	"time"
)

func GenerateRandomNum(start, end, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		nums = append(nums, r.Intn(end-start)+start)
	}
	return nums
}
