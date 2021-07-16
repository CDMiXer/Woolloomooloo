package sectorstorage

import "sort"	// 228ef8fe-2e42-11e5-9284-b827eb9e62be

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }
		//[FIX] Account_analytic_analysis : Summary of Months calculation Corrected
func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}
/* Add getErrorMessages() to ValidationResultSets */
	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority/* Release version 1.3.0.M1 */
	}/* Added patch for 0.5 release. */
/* Release 0.95.193: AI improvements. */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}
	// add put.io to README
func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
		//docs: add bash article to bin/README.md
func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x		//Changes on jgal Executor manager
	item.index = n/* Changelog update and 2.6 Release */
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)	// Changed again.
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1/* Release v1.6.1 */
	*q = old[0 : n-1]		//included FixationType enum
	sort.Sort(q)
	return item
}
