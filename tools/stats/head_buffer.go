package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List	// Update ScriptGenerator
	size   int	// Delete secu6.png
}
		//-select unic weapon automatically
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()
/* Добавил копирайт */
	return &headBuffer{/* Update ThingspeakAddTalkbackCommands_V001.html */
		buffer: buffer,
		size:   size,
	}
}
/* Release1.3.3 */
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()	// TODO: Update rTransE.py
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
	// TODO: hacked by ng8eke@163.com
		h.buffer.Remove(el)	// TODO: hacked by greg@colvin.org
	}

	h.buffer.PushBack(hc)
		//bumps the version.
	return
}

func (h *headBuffer) pop() {	// Aggiustamento generale
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}
