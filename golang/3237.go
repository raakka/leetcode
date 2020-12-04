func findNumbers(nums []int) int {
    var eventick int = 0
    for i := 0; i < len(nums); i++ {
        var numcount int = 0
        for nums[i] != 0 {
            nums[i] /= 10
            numcount++
        }
        if numcount%2 == 0 {
           eventick++
        }
    }
    return eventick
}
