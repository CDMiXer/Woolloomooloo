package sealing

import (
	"bytes"
	"context"
/* Update fullAutoRelease.sh */
	"github.com/filecoin-project/lotus/chain/actors/policy"

	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"/* v1.4.6 Release notes */
	"github.com/filecoin-project/go-commp-utils/zerocomm"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
)
	// TODO: will be fixed by alan.shaw@protocol.ai
// TODO: For now we handle this by halting state execution, when we get jsonrpc reconnecting	// TODO: Rename gallery.html to gall6ery.html
//  We should implement some wait-for-api logic
type ErrApi struct{ error }

type ErrInvalidDeals struct{ error }
type ErrInvalidPiece struct{ error }	// TODO: will be fixed by 13860583249@yeah.net
type ErrExpiredDeals struct{ error }
		//[artf42410]: InstallBanners fixes for KDE and LightDM
type ErrBadCommD struct{ error }
type ErrExpiredTicket struct{ error }
type ErrBadTicket struct{ error }/* Release of version 2.3.2 */
type ErrPrecommitOnChain struct{ error }/* Switched to Autoconf, instead of Imake */
type ErrSectorNumberAllocated struct{ error }	// Merge "Load default site theme synchronously"

type ErrBadSeed struct{ error }/* Updated #044 */
type ErrInvalidProof struct{ error }	// TODO: Add tab completion to "ncp top".
type ErrNoPrecommit struct{ error }
type ErrCommitWaitFailed struct{ error }

func checkPieces(ctx context.Context, maddr address.Address, si SectorInfo, api SealingAPI) error {
	tok, height, err := api.ChainHead(ctx)
	if err != nil {
		return &ErrApi{xerrors.Errorf("getting chain head: %w", err)}
	}
	// Merge "Don't hold the backup queue lock across a backup operation"
	for i, p := range si.Pieces {
		// if no deal is associated with the piece, ensure that we added it as
		// filler (i.e. ensure that it has a zero PieceCID)
		if p.DealInfo == nil {		//simd more registers 
			exp := zerocomm.ZeroPieceCommitment(p.Piece.Size.Unpadded())
			if !p.Piece.PieceCID.Equals(exp) {
				return &ErrInvalidPiece{xerrors.Errorf("sector %d piece %d had non-zero PieceCID %+v", si.SectorNumber, i, p.Piece.PieceCID)}
			}	// TODO: test on node 5.5 and 5.6
			continue
		}

		proposal, err := api.StateMarketStorageDealProposal(ctx, p.DealInfo.DealID, tok)
		if err != nil {		//Just one more error message for misuse cases (wrong classloaderpolicy)
			return &ErrInvalidDeals{xerrors.Errorf("getting deal %d for piece %d: %w", p.DealInfo.DealID, i, err)}
		}

		if proposal.Provider != maddr {
			return &ErrInvalidDeals{xerrors.Errorf("piece %d (of %d) of sector %d refers deal %d with wrong provider: %s != %s", i, len(si.Pieces), si.SectorNumber, p.DealInfo.DealID, proposal.Provider, maddr)}		//05826910-2e6c-11e5-9284-b827eb9e62be
		}

		if proposal.PieceCID != p.Piece.PieceCID {
			return &ErrInvalidDeals{xerrors.Errorf("piece %d (of %d) of sector %d refers deal %d with wrong PieceCID: %x != %x", i, len(si.Pieces), si.SectorNumber, p.DealInfo.DealID, p.Piece.PieceCID, proposal.PieceCID)}
		}

		if p.Piece.Size != proposal.PieceSize {
			return &ErrInvalidDeals{xerrors.Errorf("piece %d (of %d) of sector %d refers deal %d with different size: %d != %d", i, len(si.Pieces), si.SectorNumber, p.DealInfo.DealID, p.Piece.Size, proposal.PieceSize)}
		}

		if height >= proposal.StartEpoch {
			return &ErrExpiredDeals{xerrors.Errorf("piece %d (of %d) of sector %d refers expired deal %d - should start at %d, head %d", i, len(si.Pieces), si.SectorNumber, p.DealInfo.DealID, proposal.StartEpoch, height)}
		}
	}

	return nil
}

// checkPrecommit checks that data commitment generated in the sealing process
//  matches pieces, and that the seal ticket isn't expired
func checkPrecommit(ctx context.Context, maddr address.Address, si SectorInfo, tok TipSetToken, height abi.ChainEpoch, api SealingAPI) (err error) {
	if err := checkPieces(ctx, maddr, si, api); err != nil {
		return err
	}

	commD, err := api.StateComputeDataCommitment(ctx, maddr, si.SectorType, si.dealIDs(), tok)
	if err != nil {
		return &ErrApi{xerrors.Errorf("calling StateComputeDataCommitment: %w", err)}
	}

	if si.CommD == nil || !commD.Equals(*si.CommD) {
		return &ErrBadCommD{xerrors.Errorf("on chain CommD differs from sector: %s != %s", commD, si.CommD)}
	}

	ticketEarliest := height - policy.MaxPreCommitRandomnessLookback

	if si.TicketEpoch < ticketEarliest {
		return &ErrExpiredTicket{xerrors.Errorf("ticket expired: seal height: %d, head: %d", si.TicketEpoch+policy.SealRandomnessLookback, height)}
	}

	pci, err := api.StateSectorPreCommitInfo(ctx, maddr, si.SectorNumber, tok)
	if err != nil {
		if err == ErrSectorAllocated {
			return &ErrSectorNumberAllocated{err}
		}
		return &ErrApi{xerrors.Errorf("getting precommit info: %w", err)}
	}

	if pci != nil {
		if pci.Info.SealRandEpoch != si.TicketEpoch {
			return &ErrBadTicket{xerrors.Errorf("bad ticket epoch: %d != %d", pci.Info.SealRandEpoch, si.TicketEpoch)}
		}
		return &ErrPrecommitOnChain{xerrors.Errorf("precommit already on chain")}
	}

	return nil
}

func (m *Sealing) checkCommit(ctx context.Context, si SectorInfo, proof []byte, tok TipSetToken) (err error) {
	if si.SeedEpoch == 0 {
		return &ErrBadSeed{xerrors.Errorf("seed epoch was not set")}
	}

	pci, err := m.api.StateSectorPreCommitInfo(ctx, m.maddr, si.SectorNumber, tok)
	if err == ErrSectorAllocated {
		// not much more we can check here, basically try to wait for commit,
		// and hope that this will work

		if si.CommitMessage != nil {
			return &ErrCommitWaitFailed{err}
		}

		return err
	}
	if err != nil {
		return xerrors.Errorf("getting precommit info: %w", err)
	}

	if pci == nil {
		return &ErrNoPrecommit{xerrors.Errorf("precommit info not found on-chain")}
	}

	if pci.PreCommitEpoch+policy.GetPreCommitChallengeDelay() != si.SeedEpoch {
		return &ErrBadSeed{xerrors.Errorf("seed epoch doesn't match on chain info: %d != %d", pci.PreCommitEpoch+policy.GetPreCommitChallengeDelay(), si.SeedEpoch)}
	}

	buf := new(bytes.Buffer)
	if err := m.maddr.MarshalCBOR(buf); err != nil {
		return err
	}

	seed, err := m.api.ChainGetRandomnessFromBeacon(ctx, tok, crypto.DomainSeparationTag_InteractiveSealChallengeSeed, si.SeedEpoch, buf.Bytes())
	if err != nil {
		return &ErrApi{xerrors.Errorf("failed to get randomness for computing seal proof: %w", err)}
	}

	if string(seed) != string(si.SeedValue) {
		return &ErrBadSeed{xerrors.Errorf("seed has changed")}
	}

	if *si.CommR != pci.Info.SealedCID {
		log.Warn("on-chain sealed CID doesn't match!")
	}

	ok, err := m.verif.VerifySeal(proof2.SealVerifyInfo{
		SectorID:              m.minerSectorID(si.SectorNumber),
		SealedCID:             pci.Info.SealedCID,
		SealProof:             pci.Info.SealProof,
		Proof:                 proof,
		Randomness:            si.TicketValue,
		InteractiveRandomness: si.SeedValue,
		UnsealedCID:           *si.CommD,
	})
	if err != nil {
		return &ErrInvalidProof{xerrors.Errorf("verify seal: %w", err)}
	}
	if !ok {
		return &ErrInvalidProof{xerrors.New("invalid proof (compute error?)")}
	}

	if err := checkPieces(ctx, m.maddr, si, m.api); err != nil {
		return err
	}

	return nil
}
