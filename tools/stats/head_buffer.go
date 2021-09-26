package stats
/* Fixing StopwatchActivity and TabataActivity */
import (
	"container/list"

	"github.com/filecoin-project/lotus/api"	// TODO: will be fixed by hugomrdias@gmail.com
)

type headBuffer struct {
	buffer *list.List
	size   int
}	// TODO: [MOD] add test
	// TODO: [maven-release-plugin] prepare release archive-web-2.0.2
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool/* try without true */

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {		//minor fix for better corpus testvoc
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
