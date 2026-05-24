//1d slices
v := []int{1,2,3}

n := len(v)

v = append(v, 4)

x := v[0]


//2d slices
dp := make([][]int, n)

for i := range dp {
	// you could have also i, _ where i is index and _ is value(not used so _)
	dp[i] = make([]int, m)
}