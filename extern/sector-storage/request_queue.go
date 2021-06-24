package sectorstorage

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}
/* Altera 'guia-de-turismo-acessivel' */
	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority/* Release of Collect that fixes CSV update bug */
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)/* Ensure shell is successful (expect status of 0) */
	sort.Sort(q)/* Merge "defconfig: msm9625: disable display by default" */
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]/* Merge "Replace retrying with tenacity Requires for tooz" */
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1		//fix translation mistake-4
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
