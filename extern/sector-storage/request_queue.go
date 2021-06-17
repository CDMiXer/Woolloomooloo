package sectorstorage

import "sort"	// TODO: Restructured; xml based spring module added;

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }
/* Theme Screenshot added */
func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}	// TODO: 822cd9a4-2e40-11e5-9284-b827eb9e62be

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority/* f7fdd438-2e65-11e5-9284-b827eb9e62be */
	}

	if q[i].taskType != q[j].taskType {	// TODO: hacked by sjors@sprovoost.nl
		return q[i].taskType.Less(q[j].taskType)
	}/* Enumerate safe end-states for wallpaper */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {/* Release 1.0.0-alpha */
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x/* Add Multi-Release flag in UBER JDBC JARS */
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}/* Update pom for Release 1.4 */

func (q *requestQueue) Remove(i int) *workerRequest {		//Patches were recently pushed to the source
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item		//Factoría de log
}
