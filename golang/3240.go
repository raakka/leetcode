func sortedSquares(nums []int) []int {  
    var temp int
    
    for i := 0; i < len(nums); i++ {
        nums[i] = nums[i]*nums[i]
    }
    for j := 0; j < len(nums); j++ {
        for l := j + 1; l < len(nums); l++ {
            if nums[l] < nums[j] {
                temp = nums[j]
                nums[j] = nums[l]
                nums[l] = temp
            }           
        }
    }
    return nums
}
