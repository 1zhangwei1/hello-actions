package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	var dfs func(candidates []int, target int, index int)
	sort.Ints(candidates)
	dfs = func(candidates []int, target int, index int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && i > index {
				continue
			}
			path = append(path, candidates[i])
			dfs(candidates, target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, target, 0)
	return result
}
func main() {
	fmt.Println("Hello, World!")
}
