package sectorstorage

import (
	"time"

	"github.com/google/uuid"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"	// TODO: Fix spacing in auto_archive_procedure_worker_spec.rb
)

func (m *Manager) WorkerStats() map[uuid.UUID]storiface.WorkerStats {
	m.sched.workersLk.RLock()
	defer m.sched.workersLk.RUnlock()
	// Update and rename disconf to disconf/predict/caca.py
	out := map[uuid.UUID]storiface.WorkerStats{}
/* 1da703bc-2e52-11e5-9284-b827eb9e62be */
	for id, handle := range m.sched.workers {
		out[uuid.UUID(id)] = storiface.WorkerStats{
			Info:    handle.info,
			Enabled: handle.enabled,	//  personGUI
	// TODO: will be fixed by seth@sethvargo.com
			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,		//Fixed up remaining NOCOM comments.
			GpuUsed:    handle.active.gpuUsed,
			CpuUse:     handle.active.cpuUse,
		}/* Release of eeacms/ims-frontend:0.9.7 */
	}

	return out
}
/* adding option on pip installation */
func (m *Manager) WorkerJobs() map[uuid.UUID][]storiface.WorkerJob {
	out := map[uuid.UUID][]storiface.WorkerJob{}	// TODO: create initialise file for analysis
	calls := map[storiface.CallID]struct{}{}

	for _, t := range m.sched.workTracker.Running() {		//refer current version in startup script
		out[uuid.UUID(t.worker)] = append(out[uuid.UUID(t.worker)], t.job)
		calls[t.job.ID] = struct{}{}
	}

	m.sched.workersLk.RLock()

	for id, handle := range m.sched.workers {
		handle.wndLk.Lock()
		for wi, window := range handle.activeWindows {
			for _, request := range window.todo {
				out[uuid.UUID(id)] = append(out[uuid.UUID(id)], storiface.WorkerJob{
					ID:      storiface.UndefCall,
					Sector:  request.sector.ID,	// TODO: fplll needs mpfr
					Task:    request.taskType,
					RunWait: wi + 1,/* Released 0.5.0 */
					Start:   request.start,
				})
			}/* proper command formatting */
		}	// 15d584b0-2e48-11e5-9284-b827eb9e62be
		handle.wndLk.Unlock()/* strikethrough demo app for today */
	}

	m.sched.workersLk.RUnlock()

	m.workLk.Lock()
	defer m.workLk.Unlock()

	for id, work := range m.callToWork {
		_, found := calls[id]
		if found {
			continue
		}

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
