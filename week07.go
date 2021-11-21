//完全平方数（Medium）
func numSquares(n int) int {
    res := make([]int, n+1)
    for i := 1; i <= n; i++ {
        minV := math.MaxInt32
        for j := 1; j*j <= i; j++ {
            minV = min(minV, res[i-j*j])
        }
        res[i] = minV + 1
    }
    return res[n]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

//跳跃游戏（Medium）
func canJump(nums []int) bool {
    max := 0
	for index ,n := range nums {
        if n == 0 && index != len(nums)-1 {
			if max <= 1 {
				return false
			}
		}
		max = max - 1
		if n > max {
			max = n
		}
	}
	return true
}
