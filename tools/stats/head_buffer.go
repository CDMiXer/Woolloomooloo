package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"	// Audio: start/stop on UI thread
)

type headBuffer struct {
	buffer *list.List
	size   int
}/* Fix the screenshot image */
/* issue 42 : ensure runtime type of variable definition is kept */
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()	// TODO: Add api method to delete a measurement.
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)/* [deploy] Release 1.0.2 on eclipse update site */
	}/* Update index.liquid */

	h.buffer.PushBack(hc)/* fix bad line */

	return
}
		//Added entry points
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)	// TODO: will be fixed by fjl@ethereum.org
	}
}		//docs: tweak typography
