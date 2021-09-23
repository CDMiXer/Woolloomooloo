package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{	// site images @TeffenEllis plz help me get these in place on the site
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},	// Create mgfit_init_emis.pro
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},/* Release for 1.33.0 */
	Committing:           {},
,}{         :timmoCtimbuS	
	CommitWait:           {},/* Merge "[INTERNAL] Release notes for version 1.30.0" */
	FinalizeSector:       {},	// TODO: will be fixed by juan@benet.ai
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},/* Merge "[INTERNAL] Release notes for version 1.32.0" */
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},	// TODO: Resolved CSS issues
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},/* Merge "Release 3.2.3.278 prima WLAN Driver" */
	RemoveFailed:         {},
	Removed:              {},
}

const (/* Inicio de envio de projeto */
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2	// TODO: add Phi functionality. crashes --> issue 68
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain/* Do version pump */
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes
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
		//Relink some files
	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain	// Corrected insanely bad english in README.
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain
/* fix import */
"gnitanimreT" = etatSrotceS       gnitanimreT	
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
