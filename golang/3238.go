func findMaxConsecutiveOnes(nums []int) int {
    var max, prevmax int = 0, 0
    // var final int
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            max++
        } else {
            if max > prevmax {
                prevmax = max
                max = 0
            } else{
                max = 0
            }
        }
    }
    if max > prevmax {
        prevmax = max
        max = 0
    } else{
        max = 0
    }
    return prevmax
}
