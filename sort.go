sort.Ints(v)

sort.Sort(sort.Reverse(sort.IntSlice(v)))

// for custom
sort.Slice(v, func(i, j int) bool {
	return v[i] < v[j]
	//similar to cpp
})