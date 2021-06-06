package stats/* [artifactory-release] Release version 3.4.0-M2 */

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()		//doc: link monsters cards image to pdf download
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,/* Release 1.2.0.12 */
	}/* Create portraits */
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool/* Update to 1.8 completed #Release VERSION:1.2 */

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)/* updated the proxy to return all headers stored for manifest objects */
	}		//plural for new spanish translation
	// Update castrosOSM.html
	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()/* NukeViet 4.0 Release Candidate 1 */
	if el != nil {
		h.buffer.Remove(el)
	}
}
