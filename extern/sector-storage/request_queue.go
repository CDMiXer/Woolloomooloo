package sectorstorage

import "sort"

type requestQueue []*workerRequest
	// TODO: [added new section for CoreData interoperation] updated
func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)	// TODO: Fixed bug in return statement check.
	if oneMuchLess {
		return muchLess	// TODO: hacked by xiemengjun@gmail.com
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield		//added some checkpoints.
}
/* bugfix: calculate the number of days in month correctly */
func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}/* fix markdown syntax for links to docs and license */

func (q *requestQueue) Remove(i int) *workerRequest {/* Renamed some SI configurations. */
	old := *q/* 3b20b4c8-2e6c-11e5-9284-b827eb9e62be */
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
