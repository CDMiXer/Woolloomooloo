package stats
/* Initial Release (0.1) */
import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List/* default arguments */
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {/* Task #3241: Merge of latest changes in LOFAR-Release-0_96 into trunk */
		var ok bool

		el := h.buffer.Front()/* Update mavenAutoRelease.sh */
		rethc, ok = el.Value.(*api.HeadChange)/* @Release [io7m-jcanephora-0.27.0] */
		if !ok {
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
}/* Added RelatedAlbum.getReleaseDate Support */
