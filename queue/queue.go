package queue

//简单队列对象FIFO
type Queue struct {
	datas []string
}

//添加队列元素
func Put(q *Queue, v string) {
	q.datas = append(q.datas, v)
}

//获取列头元素
func Take(q *Queue) (r string) {
	if !Empty(q) {
		l := len(q.datas)
		r = q.datas[0]
		q.datas = append(q.datas[1:l])
	}
	return r
}

//队列是否为空
func Empty(q *Queue) bool {
	return len(q.datas) == 0
}

//新建先进先出队列
func NewQueue() *Queue {
	return &Queue{}
}
