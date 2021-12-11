//滑动窗口最大值（Hard）
func maxSlidingWindow(nums []int, k int) []int {
    q := []int{}
    push := func(i int) {
        for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }

    for i := 0; i < k; i++ {
        push(i)
    }

    n := len(nums)
    ans := make([]int, 1, n-k+1)
    ans[0] = nums[q[0]]
    for i := k; i < n; i++ {
        push(i)
        for q[0] <= i-k {
            q = q[1:]
        }
        ans = append(ans, nums[q[0]])
    }
    return ans
}


//掉落的方块（Hard）
func fallingSquares(positions [][]int) []int {
	n := len(positions)
	ret := make([]int, n)

	a := [][]int{}
	a = append(a, []int{0, int(1e9), 0})
	a, ret[0] = merge(a, positions[0])
	h := 0
	for i := 1; i < n; i++ {
		a, h = merge(a, positions[i])
		ret[i] = max(ret[i-1], h)
	}
	return ret
}

func merge(a [][]int, block []int) ([][]int, int) {
	left, width, right := block[0], block[1], block[0]+block[1]
	ret := [][]int{}
	i := 0
	//保留左侧所有区间
	for i < len(a) && a[i][1] <= left {
		ret = append(ret, a[i])
		i++
	}
	//第i个区间的后半段被 block 命中，先保留前半段
	ret = append(ret, []int{a[i][0], left, a[i][2]})

	// 所有被命中的区间，取最高，并且被完全命中的，无需保留
	h := 0
	for i < len(a) && a[i][1] <= right {
		h = max(h, a[i][2])
		i++
	}
	//第i个区间的后半段超出了 block
	//检查其前半段是否被命中
	if a[i][0] < right {
		h = max(h, a[i][2])
		ret = append(ret, []int{left, right, h + width})
		ret = append(ret, []int{right, a[i][1], a[i][2]}) / 切掉被命中的前半段，只保留后半段
		i++
	} else {
		ret = append(ret, []int{left, right, h + width})
	}
	//保留右侧所有区间
	for i < len(a) {
		ret = append(ret, a[i])
		i++
	}
	return ret, h + width
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
