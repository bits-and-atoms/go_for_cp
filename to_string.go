import "strconv"

//integer to string
x := 42
s := strconv.Itoa(x) // "42" tc = o(len(s)) = o(log10(x))

//int64 to str
x := int64(42)
s := strconv.FormatInt(x, 10) // base 10

//uint64 to str
s := strconv.FormatUint(x, 10)

//float to str
f := 3.14159
s := strconv.FormatFloat(f, 'f', -1, 64)
// here first f = float value
//last 64 = float64, if 32 its float32
// second 'f' is output format
//'f' // decimal
// 'e' // scientific notation
// 'E' // scientific notation with E
// 'g' // compact format
// 'G' // compact format with E

//-1 is precision, how many digits after decimal
//-1 is special prec as it means , Use the smallest number of digits necessary so that converting back gives the same float.