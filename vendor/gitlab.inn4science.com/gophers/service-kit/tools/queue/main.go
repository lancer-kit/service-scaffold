package queue

var sq *Queue

func initQueue() {
	if sq != nil {
		return
	}

	sq = sq.Init()
	return
}

func NewQueue() (q *Queue) {
	return q.Init()
}

func Put(uid string, data interface{}) {
	initQueue()
	sq.Put(uid, data)
}

func Delete(uid string) {
	initQueue()
	sq.Delete(uid)
}

func GetFirst() interface{} {
	initQueue()
	return sq.GetFirst()
}

func GetLast() interface{} {
	initQueue()
	return sq.GetLast()
}

func GetByID(uid string) interface{} {
	initQueue()
	return sq.GetByID(uid)
}

func IsInQueue(uid string) bool {
	initQueue()
	return sq.IsInQueue(uid)
}
