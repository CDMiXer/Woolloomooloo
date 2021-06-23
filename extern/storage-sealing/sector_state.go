package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{		//Merge "Fixed SiteArray serialization"
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},	// TODO: hacked by nagydani@epointsystem.org
	AddPiece:             {},
	AddPieceFailed:       {},		//build 19 - better fetch n carry
	GetTicket:            {},
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
	FailedUnrecoverable:  {},/* Preparing WIP-Release v0.1.35-alpha-build-00 */
	SealPreCommit1Failed: {},	// Add address info to organization mapper
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},	// TODO: will be fixed by seth@sethvargo.com
	ComputeProofFailed:   {},	// TODO: hacked by arachnid@notdot.net
	CommitFailed:         {},
	PackingFailed:        {},	// TODO: Agregar todas las opciones posibles jekyll-assets
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},		//Rebuilt index with cmgonza
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},
}

const (
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket/* Release as version 3.0.0 */
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed/* adding 'gis' */
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain/* Add unfinished dogleg nonlinear minimizer (not in build system yet). */
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"	// TODO: will be fixed by peterke@gmail.com
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"	// TODO: will be fixed by boringland@protonmail.ch
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"	// TODO: hacked by caojiaoyue@protonmail.com
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"	// TODO: rename to gridstack.scss
	DealsExpired         SectorState = "DealsExpired"
	RecoverDealIDs       SectorState = "RecoverDealIDs"

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
