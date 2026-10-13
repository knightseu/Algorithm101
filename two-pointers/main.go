package main

func moveZeroes(nums []int) {
	l := 0
	r := 0
	for l < len(nums) && r < len(nums) {
		for r < len(nums) && nums[r] == 0 {
			r++
		}
		if r >= len(nums) {
			continue
		}
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
		l++
		r++
	}

}
