package sectorstorage
/* Vorbereitung für Release 3.3.0 */
import "sort"
/* 0.1.0 Release Candidate 14 solves a critical bug */
type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess		//Added example config and added links to external modules.
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
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}	// TODO: Apply last changes on config.

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)/* [deploy] Release 1.0.2 on eclipse update site */
}/* Added official changelog */

func (q *requestQueue) Remove(i int) *workerRequest {	// TODO: will be fixed by alex.gaynor@gmail.com
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1		//Fix regex so .+(r+)rand doesn't match
	*q = old[0 : n-1]
	sort.Sort(q)	// TODO: custom link with analytics
	return item
}
