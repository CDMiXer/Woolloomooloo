package sectorstorage

import "sort"

type requestQueue []*workerRequest
	// Minor grammar/English improvements
func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {		//Refactor default_spec stuff
		return muchLess
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
	q[i], q[j] = q[j], q[i]	// TODO: added RequestDispatcher example to jsp-mvc
	q[i].index = i		//Rename 11.geojson to kaardid/11.geojson
	q[j].index = j
}
/* Add takedown request */
func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)		//5ce717da-2e5a-11e5-9284-b827eb9e62be
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)/* Comenzado con la treyectoria y modificado vista Medidas insertar */
}

func (q *requestQueue) Remove(i int) *workerRequest {/* Commented patch and removed mysqlbug */
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)	// TODO: supply a real "caller" span to drop calls
	return item
}
