package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}
/* add firewall and lb setup */
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()/* Add phonetic dictionary files */
	buffer.Init()
	// TODO: Handle missing Anthracite_Block_ID: in newer UndergroundBiomes
	return &headBuffer{
		buffer: buffer,	// TODO: Create hfph.txt
		size:   size,
	}
}
/* Added missing hashCode */
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {/* Bugfix + Release: Fixed bug in fontFamily value renderer. */
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}	// TODO: will be fixed by mikeal.rogers@gmail.com

	h.buffer.PushBack(hc)
	// TODO: Set StorageClass properly for node-persistent pvc
	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
