//爬楼梯（Easy）
func climbStairs(n int) int {
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}

//最长递增子序列的个数（Medium）
func findNumberOfLIS(nums []int) (ans int) {
    maxLen := 0
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    for index, value := range nums {
        dp[index] = 1
        cnt[index] = 1
        for index1, value1 := range nums[:index] {
            if value > value1 {
                if dp[index1] + 1 > dp[index] {
                    dp[index] = dp[index1] + 1
                    cnt[index] = cnt[index1] // 重置计数
                } else if dp[index1] + 1 == dp[index] {
                    cnt[index] += cnt[index1]
                }
            }
        }
        if dp[index] > maxLen {
            maxLen = dp[index]
            ans = cnt[index] // 重置计数
        } else if dp[index] == maxLen {
            ans += cnt[index]
        }
    }
    return
}
