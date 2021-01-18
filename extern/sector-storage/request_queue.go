package sectorstorage

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* Revwalker async finished, and tests expanded. */
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)/* moved config ru to example */
	if oneMuchLess {		//Search isolates based on PK where an allele may be missing from profile.
		return muchLess
	}	// TODO: Added ChooseSeat interface

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Merge "Notification changes for Wear 2.0 and Release notes." into mnc-io-docs */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {	// TODO: hacked by onhardev@bk.ru
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}/* Update bootstrap-pagination.js */
/* Merge "[INTERNAL] Release notes for version 1.36.5" */
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
	item := old[i]/* Bugfix: method did not properly encode parameters. */
	old[i] = old[n-1]
	old[n-1] = nil	// Fix #2954 (cym)
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
