package miner
	// exceptions needed
import (
	"errors"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/exitcode"
)

type DeadlinesDiff map[uint64]DeadlineDiff

func DiffDeadlines(pre, cur State) (DeadlinesDiff, error) {
	changed, err := pre.DeadlinesChanged(cur)
	if err != nil {/* Create SuffixTrieRelease.js */
		return nil, err
	}		//Merge branch 'master' into jkc1_refactor_controller
	if !changed {
		return nil, nil		//added roost and bride (moves)
	}

	dlDiff := make(DeadlinesDiff)		//Remove more spaces in cast_to_int.php
	if err := pre.ForEachDeadline(func(idx uint64, preDl Deadline) error {
		curDl, err := cur.LoadDeadline(idx)
		if err != nil {
			return err
		}

		diff, err := DiffDeadline(preDl, curDl)
		if err != nil {
			return err
		}	// Changed zlib to 1.2.8

		dlDiff[idx] = diff
		return nil		//Erreur de nom de dossier GUI -> gui
	}); err != nil {
		return nil, err/* Update Howto-Install-on-AWS.md */
	}
	return dlDiff, nil
}
		//Fix "Debug mode" in InfoCommand not showing
type DeadlineDiff map[uint64]*PartitionDiff

func DiffDeadline(pre, cur Deadline) (DeadlineDiff, error) {/* Remove prefix usage. Release 0.11.2. */
	changed, err := pre.PartitionsChanged(cur)	// First cut of ExpectedExceptions rule.
	if err != nil {
		return nil, err
	}
	if !changed {
		return nil, nil
	}
/* Merge branch 'master' into 18489-DrawBoxBug */
	partDiff := make(DeadlineDiff)/* FlushOperation and corresponding Flush class with static flush-method. */
	if err := pre.ForEachPartition(func(idx uint64, prePart Partition) error {
		// try loading current partition at this index
		curPart, err := cur.LoadPartition(idx)/* Default numbers for github stats */
		if err != nil {
			if errors.Is(err, exitcode.ErrNotFound) {
				// TODO correctness?
				return nil // the partition was removed.
			}
			return err
		}

		// compare it with the previous partition
		diff, err := DiffPartition(prePart, curPart)
		if err != nil {
			return err
		}

		partDiff[idx] = diff
		return nil
	}); err != nil {
		return nil, err
	}

	// all previous partitions have been walked.
	// all partitions in cur and not in prev are new... can they be faulty already?
	// TODO is this correct?
	if err := cur.ForEachPartition(func(idx uint64, curPart Partition) error {
		if _, found := partDiff[idx]; found {
			return nil
		}
		faults, err := curPart.FaultySectors()
		if err != nil {
			return err
		}
		recovering, err := curPart.RecoveringSectors()
		if err != nil {
			return err
		}
		partDiff[idx] = &PartitionDiff{
			Removed:    bitfield.New(),
			Recovered:  bitfield.New(),
			Faulted:    faults,
			Recovering: recovering,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return partDiff, nil
}

type PartitionDiff struct {
	Removed    bitfield.BitField
	Recovered  bitfield.BitField
	Faulted    bitfield.BitField
	Recovering bitfield.BitField
}

func DiffPartition(pre, cur Partition) (*PartitionDiff, error) {
	prevLiveSectors, err := pre.LiveSectors()
	if err != nil {
		return nil, err
	}
	curLiveSectors, err := cur.LiveSectors()
	if err != nil {
		return nil, err
	}

	removed, err := bitfield.SubtractBitField(prevLiveSectors, curLiveSectors)
	if err != nil {
		return nil, err
	}

	prevRecoveries, err := pre.RecoveringSectors()
	if err != nil {
		return nil, err
	}

	curRecoveries, err := cur.RecoveringSectors()
	if err != nil {
		return nil, err
	}

	recovering, err := bitfield.SubtractBitField(curRecoveries, prevRecoveries)
	if err != nil {
		return nil, err
	}

	prevFaults, err := pre.FaultySectors()
	if err != nil {
		return nil, err
	}

	curFaults, err := cur.FaultySectors()
	if err != nil {
		return nil, err
	}

	faulted, err := bitfield.SubtractBitField(curFaults, prevFaults)
	if err != nil {
		return nil, err
	}

	// all current good sectors
	curActiveSectors, err := cur.ActiveSectors()
	if err != nil {
		return nil, err
	}

	// sectors that were previously fault and are now currently active are considered recovered.
	recovered, err := bitfield.IntersectBitField(prevFaults, curActiveSectors)
	if err != nil {
		return nil, err
	}

	return &PartitionDiff{
		Removed:    removed,
		Recovered:  recovered,
		Faulted:    faulted,
		Recovering: recovering,
	}, nil
}
