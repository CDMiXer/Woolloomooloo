package sectorstorage	// TODO: hacked by arajasek94@gmail.com

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }
		//Merge "Add voting docs jobs to kuryr-tempest-plugin"
func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}
/* Merge "[Fabric] Merge interface related GDOs into one" */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Constants required across project */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i	// remove a legacy IsaacObjectType that didn't make sense any more
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)	// TODO: will be fixed by martin2cai@hotmail.com
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)/* Release 1.2.0-SNAPSHOT */
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q	// Merged #166.
	n := len(old)/* Remove redundant badge */
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item/* Release tag: 0.6.9. */
}	// TODO: hacked by ng8eke@163.com
