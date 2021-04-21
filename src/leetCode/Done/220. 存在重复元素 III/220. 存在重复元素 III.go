package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucket := map[int]int{}
	bucketSize := t + 1
	for i, num := range nums {
		bucketID := getBucketID(num, bucketSize)
		if _, ok := bucket[bucketID]; ok {
			return true
		}
		if v, ok := bucket[bucketID-1]; ok && abs(v-num) <= t {
			return true
		}
		if v, ok := bucket[bucketID+1]; ok && abs(v-num) <= t {
			return true
		}
		bucket[bucketID] = num
		if len(bucket) > k {
			toDelID := getBucketID(nums[i-k], bucketSize)
			delete(bucket, toDelID)
		}
	}
	return false
}

func getBucketID(num, bucketSize int) int {
	if num < 0 {
		return num/bucketSize - 1
	}
	return num / bucketSize
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
