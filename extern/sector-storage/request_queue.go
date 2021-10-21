package sectorstorage

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess		//[docs] make param name consistent
	}
/* blog post for steering committee */
	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Merge "mtd: ubi: Extend UBI layer debug/messaging capabilities" */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {/* Changing configure */
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j	// TODO: 0832b85c-2e5a-11e5-9284-b827eb9e62be
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n/* [#518] Release notes 1.6.14.3 */
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil	// TODO: will be fixed by martin2cai@hotmail.com
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item		//improvements to libstaff and libhours plugins, 10/6
}
