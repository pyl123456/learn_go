//冗余连接（Medium）
func findRedundantConnection(edges [][]int) []int {
    parent := make([]int, len(edges)+1)
    for i := range parent {
        parent[i] = i
    }
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(from, to int) bool {
        x, y := find(from), find(to)
        if x == y {
            return false
        }
        parent[x] = y
        return true
    }
    for _, e := range edges {
        if !union(e[0], e[1]) {
            return e
        }
    }
    return nil
}

//岛屿数量（Medium）
type UnionFindSet struct {
    grid [][]int
    count int 
}

func numIslands(grid [][]byte) int {
    ufs := NewUnionFindSet(grid)
    return ufs.count
}
func NewUnionFindSet(grid [][]byte) *UnionFindSet{
    ufs := &UnionFindSet{}
    if grid == nil {
        return nil
    }
    g := make([][]int, len(grid))
    for i:= 0; i < len(grid); i++ {
        g[i] = make([]int, len(grid[0]))
    }
    ufs.grid = g

    bias := [][]int{{0,-1}, {-1,0}}
    for i:=0; i<len(grid); i++ {
        for j := 0; j < len(grid[0]); j++{
            idx := i*len(grid[0]) + j
            if grid[i][j] != '1' {
                continue
            }
            ufs.grid[i][j] = idx+1
            ufs.count++
            for k := 0; k < len(bias); k++{
                x := i+bias[k][0]
                y := j+bias[k][1]
                if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
                    if grid[x][y] == '1' {
                        ufs.union(i,j, x,y)
                    }
                }
            }
        }
    }
    return ufs
}

func (ufs *UnionFindSet) find (x,y int) int {
    return ufs.grid[x][y]
}

func (ufs *UnionFindSet) union (a, b int, x,y int) {
    if ufs.grid[a][b] == ufs.grid[x][y] {
        return
    }
    var startX, startY, targetRank, newRank int
    if ufs.grid[a][b] < ufs.grid[x][y] {
        startX, startY, targetRank, newRank = x, y, ufs.grid[x][y], ufs.grid[a][b]
    }else{
        startX, startY, targetRank, newRank = a, b, ufs.grid[a][b], ufs.grid[x][y]
    }    
    ufs.count--
    ufs.FloodFill(startX, startY, targetRank, newRank) 
}

func (ufs *UnionFindSet) FloodFill(x,y int, targetRank int, newRank int) {
    if x <0 || x >= len(ufs.grid) || y < 0 || y >= len(ufs.grid[0]) {
        return
    }
    if ufs.grid[x][y] == newRank || ufs.grid[x][y] != targetRank {
        return
    }
    ufs.grid[x][y] = newRank
    ufs.FloodFill(x-1,y,targetRank, newRank)
    ufs.FloodFill(x+1,y,targetRank, newRank)
    ufs.FloodFill(x, y-1, targetRank, newRank)
    ufs.FloodFill(x, y+1, targetRank, newRank)
}
