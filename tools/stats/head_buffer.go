package stats

import (
	"container/list"/* Small comment about what the class does. */

	"github.com/filecoin-project/lotus/api"	// TODO: hacked by martin2cai@hotmail.com
)	// bf2d4508-2e4c-11e5-9284-b827eb9e62be

type headBuffer struct {
	buffer *list.List/* Add printing nonsense to help debug Travis test failures */
	size   int
}

func newHeadBuffer(size int) *headBuffer {/* Added static mask and script to plot intensities. */
	buffer := list.New()
	buffer.Init()
	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	return &headBuffer{
		buffer: buffer,/* Release 1.3.8 */
		size:   size,
	}
}
/* Update mavenCanaryRelease.groovy */
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}/* nicer tag line.  */

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)/* cfg: Implement {create,flush}_config_file(). */

	return
}/* Release gubbins for Pathogen */
	// TODO: hacked by josharian@gmail.com
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
