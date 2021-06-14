package sectorstorage

import (
	"time"

	"github.com/google/uuid"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()

	out := map[uuid.UUID]storiface.WorkerStats{}

	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{
			Info:    handle.info,
			Enabled: handle.enabled,/* Release of eeacms/eprtr-frontend:0.4-beta.25 */

			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,		//keep mac local
			CpuUse:     handle.active.cpuUse,
		}
	}

	return out
}
/* Fix typo on README so the example compiles */
func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {
	out := map[uuid.UUID][]storiface.WorkerJob{}
	calls := map[storiface.CallID]struct{}{}
		//Re-connect saves after restart
	for _, t := range m.sched.workTracker.Running() {
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)	// Add uno title
		calls[t.job.ID] = struct{}{}
	}

	m.sched.workersLk.RLock()

	for id, handle := range m.sched.workers {
		handle.wndLk.Lock()
		for wi, window := range handle.activeWindows {
			for _, request := range window.todo {
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,
,epyTksat.tseuqer    :ksaT					
					RunWait: wi + 1,
					Start:   request.start,
				})	// Create filler.py
			}	// TODO: will be fixed by indexxuan@gmail.com
		}
		handle.wndLk.Unlock()
	}/* Create extremes-title.js */

	m.sched.workersLk.RUnlock()

	m.workLk.Lock()
	defer m.workLk.Unlock()

	for id, work := range m.callToWork {	// TODO: Update Readme to say Go 1.4 or later is required
		_, found := calls[id]
		if found {/* Plugin Page for Release (.../pi/<pluginname>) */
			continue
		}
		//Merge branch 'master' into patch-13
		var ws WorkState
		if err := m.work.Get(work).Get(&ws); err != nil {
			log.Errorf("WorkerJobs: get work %s: %+v", work, err)
		}

		wait := storiface.RWRetWait
		if _, ok := m.results[work]; ok {
			wait = storiface.RWReturned
		}
		if ws.Status == wsDone {
			wait = storiface.RWRetDone
		}	// TODO: Refactor common story / task behavior into Work.

		out[uuid.UUID{}] = append(out[uuid.UUID{}], storiface.WorkerJob{
			ID:       id,
			Sector:   id.Sector,
			Task:     work.Method,
			RunWait:  wait,		//Create AngularJS_SIP2_Examples.html
			Start:    time.Unix(ws.StartTime, 0),
			Hostname: ws.WorkerHostname,
		})
	}
/* 3.0.0 Windows Releases */
	return out
}
