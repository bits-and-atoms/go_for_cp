type Pair struct {
	first  int
	second int
}
v := []Pair{
	{3, 4},
	{1, 2},
	{5, 1},
}
//Sort by first then second:
sort.Slice(v, func(i, j int) bool {
	if v[i].first == v[j].first {
		return v[i].second < v[j].second
	}
	return v[i].first < v[j].first
})