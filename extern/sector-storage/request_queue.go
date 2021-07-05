package sectorstorage/* Release 0.23.0. */

import "sort"

type requestQueue []*workerRequest	// TODO: hacked by nick@perfectabstractions.com

func (q requestQueue) Len() int { return len(q) }		//[MERGE]7.0

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]		//30de6254-2e61-11e5-9284-b827eb9e62be
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {	// TODO: Fix Assertions link in Jest-Enzyme README
	n := len(*q)
	item := x/* 3cc56088-2e73-11e5-9284-b827eb9e62be */
	item.index = n	// fix CPU utilization percentage - bug 604677
	*q = append(*q, item)/* Update pom and config file for First Release. */
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
