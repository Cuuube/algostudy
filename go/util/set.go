package util

type Set struct {
	data map[int]bool
}

func (s *Set) Add(v int) {
	if s.data == nil {
		s.data = make(map[int]bool)
	}
	s.data[v] = true
}

func (s *Set) Remove(v int) {
	if s.data == nil {
		s.data = make(map[int]bool)
	}
	delete(s.data, v)
}

func (s *Set) Clear() {
	s.data = make(map[int]bool)
}

func (s *Set) Length() int {
	if s.data == nil {
		s.data = make(map[int]bool)
	}
	return len(s.data)
}

func (s *Set) Keys() []int {
	ret := make([]int, 0)
	for k := range s.data {
		ret = append(ret, k)
	}
	return ret
}
