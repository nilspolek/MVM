package MVM

import "io"

type Stack struct {
	data []int
}

func (s *Stack) len() int {
	return len(s.data)
}
func (s *Stack) push(v int) {
	s.data = append(s.data, v)
}
func (s *Stack) pop() (int, error) {
	if s.len() == 0 {
		return 0, io.EOF
	}
	v := s.data[s.len()-1]
	s.data = s.data[:s.len()-1]
	return v, nil
}
