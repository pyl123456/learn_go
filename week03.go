//合并K 个升序链表（Hard） 
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    return merge(lists, 0 , len(lists) - 1)
}

func merge(lists []*ListNode, l, r int) *ListNode{
    if l == r{
        return lists[l]
    }
    if l > r{
        return nil
    }
    mid := (l + r) >> 1
    return mergeTwolists(merge(lists, l, mid), merge(lists, mid + 1, r))
}

func mergeTwolists(l1, l2 *ListNode) *ListNode {
    head := &ListNode{}
    cur := head
    // l1Ptr := l1
    // l2Ptr := l2
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }
    return head.Next
}




//全排列 II （Medium）
func permuteUnique(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	path := make([]int,0,0)
	sort.Ints(nums)
	dfs(&result, nums, visited, &path)
	return result
}
func dfs(result *[][]int, nums []int, visited []bool, path *[]int) {
	if len(*path) == len(nums) {
		temp := make([]int, len(nums), len(nums))
		copy(temp,*path)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
			continue
		}
		// 更改visited的bool值
		*path = append(*path, nums[i])
		visited[i] = true
		// 递归
		dfs(result, nums, visited, path)
		visited[i] = false
		*path=(*path)[:len(*path)-1]
	}
}
