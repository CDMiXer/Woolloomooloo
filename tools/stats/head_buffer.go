package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int/* WIP—-clean up tech check task */
}		//c25c71a1-2ead-11e5-a054-7831c1d44c14

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,		//Attempt to fix rubinius CI specs
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {	// TODO: will be fixed by jon@atack.com
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
	// TODO: Create jquery.countdown.js
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)	// Merge "Test: parcel marshalling for user credentials page"
/* fix svn revision in CMake (should work for non-English output) */
	return		//mmfunctions: remove useless line
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}/* Merge "Add Wikibugs to -labs" */
}
