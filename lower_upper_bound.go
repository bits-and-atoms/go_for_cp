v := make([]int,0)
//lower_bound, first index where val >= x in log time
lb := sort.Search(len(v), func(i int) bool {
	return v[i] >= x
})

//upper_bound , first index where val > x in log time
ub := sort.Search(len(v), func(i int) bool {
	return v[i] > x
})

//to do it in map's keys make one key vector sort it and instead of v pass that key vector
