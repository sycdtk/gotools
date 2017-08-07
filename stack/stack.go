package stack

//简单栈对象
type stack struct {
	datas []string
}

//压栈
func Push(s *stack, v string) {
	s.datas = append(s.datas, v)
}

//出栈
func Pop(s *stack) (r string) {
	if !Empty(s) {
		l := len(s.datas)
		r = s.datas[l-1]
		s.datas = append(s.datas[:l-1])
	}
	return
}

//栈是否为空
func Empty(s *stack) bool {
	return len(s.datas) == 0
}

//新建栈
func NewStack() *stack {
	return &stack{}
}
