package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)
		//Fix fatal onException 
type headBuffer struct {
	buffer *list.List
	size   int
}
/* Release v3.2.3 */
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{/* Delete alien-movies-timeline.md */
		buffer: buffer,
		size:   size,
	}
}		//a5483b96-2e3f-11e5-9284-b827eb9e62be

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {		//Merge "radio-tavarua: Add support for WCN2243 v2.1 SOC" into jb_rel_rb5_qrd
			panic("Value from list is not the correct type")
		}/* Release 1.34 */

		h.buffer.Remove(el)/* [JENKINS-64657] removed modifier from constructor */
	}	// Automatic changelog generation for PR #46793 [ci skip]

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
