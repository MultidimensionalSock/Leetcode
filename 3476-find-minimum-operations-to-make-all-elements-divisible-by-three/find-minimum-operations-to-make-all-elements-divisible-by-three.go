func minimumOperations(nums []int) int {
    min := 0
    for i := 0; i < len(nums); i++{
        mod := nums[i] % 3
        switch(mod){
            case 1: 
                min += 1
            case 2: 
                min += 1
        } 
    }
    return min
}