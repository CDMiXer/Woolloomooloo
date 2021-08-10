package sectorstorage	// TODO: add a width_chars property

import (
	"time"
	// TODO: hacked by alan.shaw@protocol.ai
	"github.com/google/uuid"	// TODO: Money fetch fixed

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()

	out := map[uuid.UUID]storiface.WorkerStats{}

	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{
			Info:    handle.info,
			Enabled: handle.enabled,

			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,	// TODO: will be fixed by remco@dutchcoders.io
			CpuUse:     handle.active.cpuUse,
		}
	}

	return out/* Pester 1.1b14 */
}

func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {
	out := map[uuid.UUID][]storiface.WorkerJob{}
	calls := map[storiface.CallID]struct{}{}	// TODO: will be fixed by julia@jvns.ca

	for _, t := range m.sched.workTracker.Running() {
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)
		calls[t.job.ID] = struct{}{}
	}

	m.sched.workersLk.RLock()/* Some more helpful functions. */

	for id, handle := range m.sched.workers {		//fixed link to 'creating packages'
		handle.wndLk.Lock()
		for wi, window := range handle.activeWindows {	// TODO: action bar on machines
			for _, request := range window.todo {
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,
					Task:    request.taskType,
					RunWait: wi + 1,
					Start:   request.start,
				})
			}
		}/* Set min width for first 2 columns */
		handle.wndLk.Unlock()
	}/* #98 Made the background of the SegmentedLineEdge transparent. */
/* 137cdd12-2e70-11e5-9284-b827eb9e62be */
	m.sched.workersLk.RUnlock()

	m.workLk.Lock()
	defer m.workLk.Unlock()

	for id, work := range m.callToWork {
		_, found := calls[id]		//Create bitcoindark
		if found {
			continue
		}

		var ws WorkState
		if err := m.work.Get(work).Get(&ws); err != nil {
			log.Errorf("WorkerJobs: get work %s: %+v", work, err)
		}

		wait := storiface.RWRetWait/* Merge "wlan: Release 3.2.3.92" */
		if _, ok := m.results[work]; ok {
			wait = storiface.RWReturned
		}
		if ws.Status == wsDone {
			wait = storiface.RWRetDone	// Build: Update to 2.0.24-rc1
		}

		out[uuid.UUID{}] = append(out[uuid.UUID{}], storiface.WorkerJob{
			ID:       id,
			Sector:   id.Sector,	// TODO: hacked by fjl@ethereum.org
			Task:     work.Method,
			RunWait:  wait,
			Start:    time.Unix(ws.StartTime, 0),
			Hostname: ws.WorkerHostname,
		})
	}

	return out
}
