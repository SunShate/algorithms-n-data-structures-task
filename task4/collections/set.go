package collections

type Set map[int]struct{} //empty structs occupy 0 memory

func (s *Set) Has(v int) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *Set) Add(v int) {
	(*s)[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete((*s), v)
}

func (s *Set) Clear() {
	*s = make(map[int]struct{})
}

func (s *Set) Size() int {
	return len(*s)
}

//optional functionalities

// AddMulti Add multiple values in the set
func (s *Set) AddMulti(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

func (s *Set) Iter() []int {
	var keys []int
	for k := range *s {
		keys = append(keys, k)
	}
	return keys
}
