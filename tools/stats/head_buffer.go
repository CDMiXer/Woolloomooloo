package stats
/* Release version 3.1.6 build 5132 */
import (
	"container/list"/* Release dhcpcd-6.7.0 */

	"github.com/filecoin-project/lotus/api"
)
/* Update to latest Tenant service */
type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {	// TODO: Add ACPI handling for power button
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}/* Release version 0.8.5 Alpha */
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
/* Save playlist state on destruction of service */
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}
/* Add form validator for icon_emoji */
func (h *headBuffer) pop() {/* Split homepage building into phases */
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
