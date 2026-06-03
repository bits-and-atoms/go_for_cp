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

//a previous and next greater element in stack and queue
for i := 0; i < n; i++ {
		pge[i] = -1
		nge[i] = n
		for len(st) > 0 && v[st[len(st)-1]] < v[i] {
			top := st[len(st)-1]
			nge[top] = i
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			pge[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	prev := n
	for len(st) > 0 {
		if prev != n && v[prev] == v[st[len(st)-1]] {
			nge[st[len(st)-1]] = prev
		}
		prev = st[len(st)-1]
		st = st[:len(st)-1]
	}
