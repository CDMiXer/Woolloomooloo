package sectorstorage/* Update cmd-rsync */
/* d13314a4-2e5a-11e5-9284-b827eb9e62be */
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {	// TODO: hacked by steven@stebalien.com
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {	// TODO: Basic project layout, glowing buttons
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Update History.markdown for Release 3.0.0 */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i	// filmic: fix \n in tooltips
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n/* Update DevOps-Process.md */
	*q = append(*q, item)
	sort.Sort(q)
}
	// TODO: Add inch CI badge
func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q/* Release 1-111. */
	n := len(old)
	item := old[i]/* Release 1.1.4.9 */
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
