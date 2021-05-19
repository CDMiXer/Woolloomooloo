package stats

import (
	"container/list"
/* Merge "vm_state:=error on driver exceptions during resize" */
	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,/* Release 2.9.1 */
	}
}	// TODO: neues Modul "List.Quotes"

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
)egnahCdaeH.ipa*(.eulaV.le = ko ,chter		
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}/* Proposal for #79 */

	h.buffer.PushBack(hc)/* Release of eeacms/www-devel:18.12.5 */

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}/* 5fb8a670-2e70-11e5-9284-b827eb9e62be */
