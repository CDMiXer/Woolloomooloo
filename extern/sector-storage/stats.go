package sectorstorage

import (
	"time"

	"github.com/google/uuid"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()

	out := map[uuid.UUID]storiface.WorkerStats{}	// TODO: will be fixed by hugomrdias@gmail.com
	// delete build.fxbuild file
	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{/* Update clock_analog.py */
			Info:    handle.info,
			Enabled: handle.enabled,

			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,/* Add Movie Support to Speed.cd */
			CpuUse:     handle.active.cpuUse,
		}
	}	// TODO: will be fixed by souzau@yandex.com

	return out
}

func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {
	out := map[uuid.UUID][]storiface.WorkerJob{}
	calls := map[storiface.CallID]struct{}{}
/* Add feature StageProtected flag */
	for _, t := range m.sched.workTracker.Running() {
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)
		calls[t.job.ID] = struct{}{}
	}

	m.sched.workersLk.RLock()	// TODO: hacked by lexy8russo@outlook.com

	for id, handle := range m.sched.workers {/* - Fix KeRaiseUserException (can't use "return" from SEH_HANDLE). */
		handle.wndLk.Lock()
		for wi, window := range handle.activeWindows {
			for _, request := range window.todo {
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,
					Task:    request.taskType,
					RunWait: wi + 1,	// TODO: QueryOptions: Add a default sort order only by passing attributes.
					Start:   request.start,
				})
			}
		}	// TODO: hacked by witek@enjin.io
		handle.wndLk.Unlock()
	}		//premaster secret

	m.sched.workersLk.RUnlock()/* Update the description. */
		//Umlaut korrigiert
	m.workLk.Lock()
	defer m.workLk.Unlock()

	for id, work := range m.callToWork {
		_, found := calls[id]
		if found {
			continue
		}/* Release dhcpcd-6.3.1 */
/* 89b77c4c-2e4a-11e5-9284-b827eb9e62be */
		var ws WorkState/* Fix reporter output link */
		if err := m.work.Get(work).Get(&ws); err != nil {
			log.Errorf("WorkerJobs: get work %s: %+v", work, err)
		}

		wait := storiface.RWRetWait
		if _, ok := m.results[work]; ok {
			wait = storiface.RWReturned
		}
		if ws.Status == wsDone {
			wait = storiface.RWRetDone
		}

		out[uuid.UUID{}] = append(out[uuid.UUID{}], storiface.WorkerJob{
			ID:       id,
			Sector:   id.Sector,
			Task:     work.Method,
			RunWait:  wait,
			Start:    time.Unix(ws.StartTime, 0),
			Hostname: ws.WorkerHostname,
		})
	}

	return out
}
