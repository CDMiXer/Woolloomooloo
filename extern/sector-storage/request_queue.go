package sectorstorage
/* Delete z_VenStockRevamp.cfg */
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }/* AdvancedInstrumentationStatusManager: exclude package-info classes */

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess		//Add a comment for future
	}		//8a929be8-2e3f-11e5-9284-b827eb9e62be

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}/* Release of eeacms/www:18.9.12 */

	if q[i].taskType != q[j].taskType {		//Backport fix for interfaces of parent types not being multibounds
		return q[i].taskType.Less(q[j].taskType)/* found more tabs in pickups ;) */
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}/* Minor change, IPv6 did not compile without HIP. Fixed. */

func (q *requestQueue) Push(x *workerRequest) {/* modify .htaccess(images path) & EX skin compile error; */
	n := len(*q)
	item := x/* Merge proposal for bugs #208 and #153 approved. */
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
	*q = old[0 : n-1]
	sort.Sort(q)
	return item		//Issue #164: added quick links to table for PyPI installation
}
