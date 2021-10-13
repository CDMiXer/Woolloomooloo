package stats		//Added new runes.

import (	// TODO: Merge "msm: camera: Adjust exposure setting" into ics_chocolate
	"container/list"
	// Rename Argonne to Argonne.md
	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {/* Merge "FAB-5989 Release Hyperledger Fabric v1.0.2" */
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

		el := h.buffer.Front()	// TODO: Merge "msm: camera: enable both rotary and S5 toggles for ADP platform"
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}/* Merge "Release stack lock after export stack" */

func (h *headBuffer) pop() {	// TODO: Tests with different ICP implementations.
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
