package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)
/* Fix for #172. */
type headBuffer struct {
	buffer *list.List
	size   int
}
/* Bump version. Release 2.2.0! */
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()	// Made locks work the way that normal humans would expect
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {/* [artifactory-release] Release version 1.3.2.RELEASE */
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)		//Delete Erde2.png
		if !ok {
			panic("Value from list is not the correct type")
		}		//Add docker hub link

		h.buffer.Remove(el)/* no farsi attr */
	}	// Merge "Remove unnecessary indentation level in $.wikibase.entityselector"

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {	// get the generators working
		h.buffer.Remove(el)
	}
}
