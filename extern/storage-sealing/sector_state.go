package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},/* #7 Branch page */
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},		//- [usp10] Fix build. Wine is out of control...
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},
	Committing:           {},
	SubmitCommit:         {},
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},/* Improvements on paginator */
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},/* Update Release Notes for 3.0b2 */
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},		//Simplified assert
	Removed:              {},		//Setminplayers.java
}

const (
	UndefinedSectorState SectorState = ""

	// happy path/* Be robust when merge_hash file_id not in inventory */
	Empty          SectorState = "Empty"         // deprecated/* Release v1.0.0 */
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain/* Update RFC0013-PowerShellGet-PowerShellGallery_PreRelease_Version_Support.md */
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain/* Release version 0.3.6 */
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"		//Updated Debian (markdown)
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"	// Delete inti.html
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"/* Release 3.2 104.10. */
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"
	RecoverDealIDs       SectorState = "RecoverDealIDs"		//rev 475097

	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain
	// TODO: will be fixed by arajasek94@gmail.com
	Terminating       SectorState = "Terminating"
	TerminateWait     SectorState = "TerminateWait"	// added test file
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
