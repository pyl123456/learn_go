//数组的度（Easy）https://leetcode-cn.com/problems/degree-of-an-array/
func findShortestSubArray(nums []int) int {
    counts := map[int][]int{}
    for i := 0; i< len(nums); i++{
        v := nums[i]
        if info, ok := counts[v]; ok{
            //更新nums[i]出现次数
            info[0] ++
            //更新nums[i]最后一次出现的位置
            info[2] = i
        }else{
            //记录nums[i]出现的次数为1，第一次和最后一次出现nums[i]的位置为i
            counts[nums[i]] = []int{1, i, i}
        }
    }
    var degreeInfo []int
    for _, value := range counts{
        if degreeInfo == nil{
            degreeInfo = value
            continue
        }
        //连续子数组长度
        deLength := degreeInfo[2] - degreeInfo[1]
        vLength := value[2] - value[1]
        if value[0] > degreeInfo[0]{ 
            degreeInfo = value
            continue
        }
        if value[0] == degreeInfo[0] && vLength < deLength{ 
            degreeInfo = value
        }

    }
    return degreeInfo[2] - degreeInfo[1] + 1
}

//和为 K 的子数组（Medium） https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
    sub := map[int]int{}
    sub[0] = 1
    count, pre := 0, 0
    for i := 0; i < len(nums); i ++{
        //前缀和
        pre += nums[i]
        if value, ok := sub[pre - k]; ok{
            count += value
        }
        sub[pre] += 1
    }
    return count
}
