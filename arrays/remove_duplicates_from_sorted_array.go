package arrays

import (
	"fmt"
)

func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ixStart := 0

	for ; ixStart < len(nums); ixStart++ {
		ixStop := ixStart
		fmt.Println("A", ixStart, ixStop)

		for ; ixStop < len(nums) && nums[ixStop] == nums[ixStart]; ixStop++ {
		}

		if ixStop > ixStart {
			copy(nums[ixStart+1:], nums[ixStop:])
		}

		if ixStop == len(nums) {
			break
		}

		fmt.Println("B", ixStart, ixStop)
	}

	return ixStart + 1
}
