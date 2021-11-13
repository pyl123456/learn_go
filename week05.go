//在 D 天内送达包裹的能力（Medium）
func shipWithinDays(weights []int, days int) int {
    // 确定二分查找左右边界
    left, right := 0, 0
    for _, weight := range weights {
        if weight > left {
            left = weight
        }
        right += weight
    }
    return left + sort.Search(right-left, func(x int) bool {
        x += left
        day := 1 
        sum := 0 
        for _, weight := range weights {
            if sum + weight > x {
                day ++
                sum = 0
            }
            sum += weight
        }
        return day <= days
    })
}

//爱吃香蕉的珂珂（Medium）
func minEatingSpeed(piles []int, H int) int {
	left, right, mid := 1, 0, 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	for left < right {
		mid = (left + right) / 2
		if canEatAll(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canEatAll(piles []int, speed, H int) bool {
	time := 0
	for _, pile := range piles {
		time += int(math.Ceil(float64(pile) / float64(speed)))
	}
	return time <= H
}
