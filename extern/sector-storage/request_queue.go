package sectorstorage

import "sort"

type requestQueue []*workerRequest
/* Fix: fullscreen window after input for win8 */
func (q requestQueue) Len() int { return len(q) }		//added server.sh script

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess	// TODO: hacked by jon@atack.com
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}	// TODO: Added DateValidation annotation

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
}	

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]		//Merge branch 'master' into bm500
	q[i].index = i		//-Removed Tup leftovers
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]		//3cf25016-2e6e-11e5-9284-b827eb9e62be
	sort.Sort(q)
	return item
}
