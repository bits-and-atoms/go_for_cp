//as strings are immutable so first use byte or rune array then copy it back
//for normal ascii chars
s := "hello"

b := []byte(s)

b[0] = 'H'

s = string(b)

fmt.Println(s)

//for unicode ones
s := "héllo"

r := []rune(s)

r[0] = 'H'

s = string(r)