
import "sort"

func maxProductDifference(nums []int) int {
 sort.Ints(nums); n := len(nums)
 return nums[n-1] * nums[n-2] - nums[1] * nums[0]
}
