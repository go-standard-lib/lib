package queue

type queueList[E any] struct {
	queue *[]E
}

func New[E any]() queueList[E] {
	return queueList[E]{
		queue: &[]E{},
	}
}

func (q queueList[E]) Add(elements ...E) {
	*q.queue = append(elements, *q.queue...)
}

func (q queueList[E]) Poll() E {
	index := len(*q.queue) - 1
	s := *q.queue
	head := s[index]
	*q.queue = append(s[:index], s[index+1:]...)
	return head
}

func (q queueList[E]) Peak() E {
	index := len(*q.queue) - 1
	s := *q.queue
	return s[index]
}

func (q queueList[E]) IsEmpty() bool {
	return len(*q.queue) == 0
}

func (q queueList[E]) Size() int {
	return len(*q.queue)
}
