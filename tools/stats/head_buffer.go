package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {/* Delete post_curiosity.jpg */
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}/* Release version 2.3 */
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()/* Route "Can i build X" queries via the appropriate ProductionQueue */
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {/* Started building infrastructure for shared_in settings. */
			panic("Value from list is not the correct type")
		}
/* db5b19d2-2f8c-11e5-9dae-34363bc765d8 */
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)		//add Target#test_connection method

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
