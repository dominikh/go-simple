package pkg

type S struct {
	M1 int
	M2 int
	S  *S
}

func fn() {
	const c = 0
	var v int
	var s = S{1, 2, nil}
	var s1 = S{1, 2, &s}
	_ = make([]int, 1)
	_ = make([]int, 0)          // MATCH /when length is zero, length can be omitted/
	_ = make([]int, c)          // MATCH /when length is zero, length can be omitted/
	_ = make([]int, 1, 1)       // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, c, c)       // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, v, v)       // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, v, v)       // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, s.M1, s.M1) // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, s.M1, s.M2)
	_ = make([]int, s1.S.M1, s1.S.M1) // MATCH /when length equals capacity, capacity can be omitted/
	_ = make([]int, s1.S.M1, s1.S.M2)
	_ = make([]int, 1, 2)
}
