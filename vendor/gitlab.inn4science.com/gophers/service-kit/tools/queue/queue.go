package queue

import (
	"container/list"
	"sync"
)

type Queue struct {
	list     *list.List
	index    map[string]struct{}
	elements map[string]*list.Element

	rwMutex sync.RWMutex
}

func (*Queue) Init() *Queue {
	q := &Queue{
		list:     list.New(),
		index:    make(map[string]struct{}),
		elements: make(map[string]*list.Element),
		rwMutex:  sync.RWMutex{},
	}
	return q
}

func (q *Queue) Put(uid string, data interface{}) {
	if q.IsInQueue(uid) {
		return
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	q.elements[uid] = q.list.PushBack(data)
	q.index[uid] = struct{}{}
}

func (q *Queue) Delete(uid string) {
	if !q.IsInQueue(uid) {
		return
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	el := q.elements[uid]
	q.list.Remove(el)

	delete(q.elements, uid)
	delete(q.index, uid)
}

func (q *Queue) GetFirst() interface{} {
	el := q.list.Front()
	if el == nil {
		return nil
	}
	return el.Value
}

func (q *Queue) GetLast() interface{} {
	el := q.list.Back()
	if el == nil {
		return nil
	}
	return el.Value
}

func (q *Queue) GetByID(uid string) interface{} {
	if !q.IsInQueue(uid) {
		return nil
	}

	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	el := q.elements[uid]
	if el == nil {
		return nil
	}
	return el.Value
}

func (q *Queue) IsInQueue(uid string) bool {
	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	_, ok := q.index[uid]
	return ok
}
