//queue
q := []int{}

q = append(q, x)

front := q[0]
if len(q) == 0{
	break
}
q = q[1:]
//this just masks the slice's view dont remove the first element
//for int type queue you dont need to remove element literally as int is small but when queue is of pointer struct or map type then
//you could first also do q[0] = nil so that garbage collector can collect it


//stack
st := []int{}

st = append(st, x)

top := st[len(st)-1]
// Essential memory cleanup for pointer/object stacks:
// st[len(st)-1] = nil
st = st[:len(st)-1]