package sectorstorage

import "sort"
/* added quantities descr, node */
type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }
	// TODO: hacked by boringland@protonmail.ch
func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {/* abstracted ReleasesAdapter */
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority/* - Fix a bug in ExReleasePushLock which broken contention checking. */
	}
/* Release version v0.2.7-rc007. */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)		//Update showSearch.html
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
	item := x/* [artifactory-release] Release version 2.2.1.RELEASE */
	item.index = n
	*q = append(*q, item)/* Release 0.37.0 */
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {	// TODO: strange code removed. Reverting to c++98.
	old := *q		//Changed Deploymeny Name
	n := len(old)		//Rename Bbcode to BBcode
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil/* add rfduino plugin */
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
