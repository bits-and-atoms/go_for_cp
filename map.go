mp := make(map[rune]int)
//rune is used for char here , its utf 8, you could also use byte
for _, ch := range s {
	mp[ch]++
	val,ok := mp[ch] // val to get value of this key and ok if key exists
}
delete(mp, 'a')

//to make ordered map, you can store keys in a vector and sort them
keys := make([]int, 0)

for k := range mp {
	keys = append(keys, k)
}

sort.Ints(keys)


