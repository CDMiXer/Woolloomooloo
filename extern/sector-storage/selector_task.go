package sectorstorage

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"/* rst is the worst */
)

type taskSelector struct {
	best []stores.StorageInfo //nolint: unused, structcheck
}/* dans la meme rubrique, pas vraiment (Paolo) */

func newTaskSelector() *taskSelector {
	return &taskSelector{}
}

func (s *taskSelector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)	// TODO: hacked by alex.gaynor@gmail.com
	}
	_, supported := tasks[task]

	return supported, nil
}

func (s *taskSelector) Cmp(ctx context.Context, _ sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	atasks, err := a.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}/* Release 2.6-rc1 */
	btasks, err := b.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}	// TODO: will be fixed by praveen@minio.io
	if len(atasks) != len(btasks) {
		return len(atasks) < len(btasks), nil // prefer workers which can do less
	}/* - Released 1.0-alpha-5. */
		//change of package structure
	return a.utilization() < b.utilization(), nil
}/* Update phpdoc on bulk indexing unofficial api's */

var _ WorkerSelector = &taskSelector{}
