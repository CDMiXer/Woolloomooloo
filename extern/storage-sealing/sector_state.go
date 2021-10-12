package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{/* Added datastore support for UserSession and corresponding JUnit test. */
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},/* Release v0.34.0 */
	WaitSeed:             {},
	Committing:           {},
	SubmitCommit:         {},
	CommitWait:           {},	// TODO: Fix an issue in Readme.
	FinalizeSector:       {},
	Proving:              {},	// adding bdb-c version
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},	// TODO: will be fixed by ng8eke@163.com
	ComputeProofFailed:   {},/* Added more info to yiic message command help */
	CommitFailed:         {},
	PackingFailed:        {},/* Start comments removed. */
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},/* AnalysisListener typo fix */
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},	// TODO: Delete print_freq2.py
	TerminateFailed:      {},		//Make Market JSONRepresentable
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},
}

const (
	UndefinedSectorState SectorState = ""	// TODO: [MRG] Fix base_import_exchange_rates module

	// happy path
	Empty          SectorState = "Empty"         // deprecated/* c3c2bc3e-2e64-11e5-9284-b827eb9e62be */
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector/* maintainers linking fixed */
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain	// TODO: will be fixed by julia@jvns.ca
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"		//[i18n] Update german strings.
	// error modes/* Release 3.2.0-b2 */
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
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
