package stats/* Release jedipus-2.6.16 */

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {/* Release 1.4.4 */
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,/* Yasuo Removed */
	}
}
	// TODO: partial matching
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool	// TODO: grunt bootstrap mkdirs task

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)/* Update tqdm from 4.19.7 to 4.19.9 */
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}
/* Merge "6.0 Release Number" */
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
