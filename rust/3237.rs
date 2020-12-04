impl Solution {
    pub fn find_numbers(nums: Vec<i32>) -> i32 {
        let mut eventick = 0;
        for mut number in nums.into_iter() {
            let mut numcount = 0;
            while number != 0 {
                number /= 10;
                numcount += 1;
            }
            if numcount%2 == 0 {
                eventick += 1;
            }
        }
        return eventick; 
    }
}
