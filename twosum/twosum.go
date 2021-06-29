package twosum

func TwoSum(nums []int, target int) []int {
	var arr []int
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := tmp[target-nums[i]]; ok && val != i {
			arr = append(arr, val)
			arr = append(arr, i)
			break
		}
		tmp[nums[i]] = i
	}
	return (arr)
}
