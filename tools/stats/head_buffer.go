package stats

import (
	"container/list"
/* Release woohoo! */
	"github.com/filecoin-project/lotus/api"
)
/* Replaced markForSerialization() method with theIsInUnsafeContext data member */
type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{	// Create OLT-22.html
		buffer: buffer,
		size:   size,
	}/* Release details for Launcher 0.44 */
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {	// TODO: hacked by remco@dutchcoders.io
		var ok bool

		el := h.buffer.Front()/* Add Release 1.1.0 */
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {/* Release v1.4 */
			panic("Value from list is not the correct type")
		}	// TODO: will be fixed by mail@overlisted.net

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return/* Corrected `crn tree` description */
}
	// TODO: hacked by steven@stebalien.com
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {	// Bugfix for generating bigwigs: Get slice lengths directly from database.
		h.buffer.Remove(el)/* Release 7.0 */
	}
}
