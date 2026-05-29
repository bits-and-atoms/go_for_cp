x, err := strconv.Atoi("123")
//tc o(len(str))
x, err := strconv.ParseInt("123", 10, 64)
//base,bitsize
f, err := strconv.ParseFloat("3.14", 64)
b, err := strconv.ParseBool("true")
