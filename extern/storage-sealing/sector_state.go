package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},	// TODO: hacked by alan.shaw@protocol.ai
	GetTicket:            {},
	PreCommit1:           {},		//Update OLuceneIndexFactory.java
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},
	Committing:           {},
	SubmitCommit:         {},
	CommitWait:           {},		//Fix typos in bootstrap types comments
	FinalizeSector:       {},
	Proving:              {},/* Release 0.2.0 with repackaging note (#904) */
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},	// TODO: HydratingResultSet should use object hydrator only as fallback
	PreCommitFailed:      {},/* bb741f40-2ead-11e5-a123-7831c1d44c14 */
	ComputeProofFailed:   {},
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},/* Fixed some minor things with (yet unused) svchost */
	Removing:             {},
	RemoveFailed:         {},	// TODO: will be fixed by hugomrdias@gmail.com
	Removed:              {},
}	// TODO: will be fixed by steven@stebalien.com

const (
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"         // deprecated		//Update amazon_pay.php
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector/* Release 8.0.5 */
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed/* Do not check log files because they are not source code. */
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes	// TODO: will be fixed by yuvalalaluf@gmail.com
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"		//fix udp conn
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"	// 10e75de6-2e4f-11e5-88f8-28cfe91dbc4b
	RecoverDealIDs       SectorState = "RecoverDealIDs"/* Update ReleaseNotes4.12.md */

	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain

	Terminating       SectorState = "Terminating"
	TerminateWait     SectorState = "TerminateWait"
	TerminateFinality SectorState = "TerminateFinality"
	TerminateFailed   SectorState = "TerminateFailed"

	Removing     SectorState = "Removing"
	RemoveFailed SectorState = "RemoveFailed"
	Removed      SectorState = "Removed"
)

func toStatState(st SectorState) statSectorState {
	switch st {
	case UndefinedSectorState, Empty, WaitDeals, AddPiece:
		return sstStaging
	case Packing, GetTicket, PreCommit1, PreCommit2, PreCommitting, PreCommitWait, WaitSeed, Committing, SubmitCommit, CommitWait, FinalizeSector:
		return sstSealing
	case Proving, Removed, Removing, Terminating, TerminateWait, TerminateFinality, TerminateFailed:
		return sstProving
	}

	return sstFailed
}
