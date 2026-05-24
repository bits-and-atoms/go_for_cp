//deep copy
a := []int{1,2,3}

b := make([]int, len(a))

copy(b, a)
//b is now a but changing b dont affect a

//shallow copy as slices are reference
a := []int{1,2,3}
b := a

b[0] = 100
// it also changes a
fmt.Println(a)
