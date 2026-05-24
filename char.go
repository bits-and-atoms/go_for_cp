//to convert rune to string
string(ch)

mp := make(map[rune]int)
// rune is used for char here , its utf 8, you could also use byte
for _, ch := range s {
	mp[ch]++
	val, ok := mp[ch] // val to get value of this key and ok if key exists
}