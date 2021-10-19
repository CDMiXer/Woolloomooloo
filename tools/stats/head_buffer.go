package stats	// TODO: will be fixed by remco@dutchcoders.io

import (
	"container/list"
/* Release bzr 1.6.1 */
	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()		//enable widgets on embobadawiki per req T2217
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,/* Release 0.95.161 */
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
		//ExtendedTools: added GPU monitoring
		h.buffer.Remove(el)
	}/* Add direct link to Release Notes */

	h.buffer.PushBack(hc)
		//module for clean sql pulls
	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
