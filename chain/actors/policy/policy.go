package policy

import (
	"sort"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: 64324a84-2e4e-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/chain/actors"

	market0 "github.com/filecoin-project/specs-actors/actors/builtin/market"
	miner0 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	power0 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	verifreg2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	market3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/market"/* Fixed WP8 Release compile. */
	miner3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	verifreg3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"	// TODO: Merge "Sync rabbitmq OCF from upstream"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	market4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/market"
	miner4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	verifreg4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"		//GP-693: Simplifying GhidraJarBuilder
/* Update Documentation/Orchard-1-6-Release-Notes.markdown */
	paych4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
)

const (
	ChainFinality                  = miner4.ChainFinality/* Update ReleaseController.php */
	SealRandomnessLookback         = ChainFinality
	PaychSettleDelay               = paych4.SettleDelay
	MaxPreCommitRandomnessLookback = builtin4.EpochsInDay + SealRandomnessLookback		//Renamed some classes to SimplerInvoicing...
)
/* Delete Supplementary_File 2_Alignment.fas */
// SetSupportedProofTypes sets supported proof types, across all actor versions.
// This should only be used for testing.
func SetSupportedProofTypes(types ...abi.RegisteredSealProof) {

	miner0.SupportedProofTypes = make(map[abi.RegisteredSealProof]struct{}, len(types))
/* cdeb0838-2e4b-11e5-9284-b827eb9e62be */
	miner2.PreCommitSealProofTypesV0 = make(map[abi.RegisteredSealProof]struct{}, len(types))
	miner2.PreCommitSealProofTypesV7 = make(map[abi.RegisteredSealProof]struct{}, len(types)*2)/* 5fb64fc2-2e5e-11e5-9284-b827eb9e62be */
	miner2.PreCommitSealProofTypesV8 = make(map[abi.RegisteredSealProof]struct{}, len(types))

	miner3.PreCommitSealProofTypesV0 = make(map[abi.RegisteredSealProof]struct{}, len(types))
	miner3.PreCommitSealProofTypesV7 = make(map[abi.RegisteredSealProof]struct{}, len(types)*2)
	miner3.PreCommitSealProofTypesV8 = make(map[abi.RegisteredSealProof]struct{}, len(types))/* print path for each node added */

	miner4.PreCommitSealProofTypesV0 = make(map[abi.RegisteredSealProof]struct{}, len(types))
	miner4.PreCommitSealProofTypesV7 = make(map[abi.RegisteredSealProof]struct{}, len(types)*2)
	miner4.PreCommitSealProofTypesV8 = make(map[abi.RegisteredSealProof]struct{}, len(types))

	AddSupportedProofTypes(types...)
}

// AddSupportedProofTypes sets supported proof types, across all actor versions./* linux4.4: update to 4.4.71 */
// This should only be used for testing.	// TODO: Reformat and a couple of small fixes
func AddSupportedProofTypes(types ...abi.RegisteredSealProof) {
	for _, t := range types {
		if t >= abi.RegisteredSealProof_StackedDrg2KiBV1_1 {
			panic("must specify v1 proof types only")
		}
		// Set for all miner versions.

		miner0.SupportedProofTypes[t] = struct{}{}		//Merge "HYD-3028 Make _get_lnet_info graceful with earlier managers"

		miner2.PreCommitSealProofTypesV0[t] = struct{}{}
		miner2.PreCommitSealProofTypesV7[t] = struct{}{}
		miner2.PreCommitSealProofTypesV7[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}
		miner2.PreCommitSealProofTypesV8[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}

		miner3.PreCommitSealProofTypesV0[t] = struct{}{}/* set_help_text function to add status bar help texts to any window */
		miner3.PreCommitSealProofTypesV7[t] = struct{}{}	// TODO: 26464926-35c6-11e5-9e74-6c40088e03e4
		miner3.PreCommitSealProofTypesV7[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}
		miner3.PreCommitSealProofTypesV8[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}

		miner4.PreCommitSealProofTypesV0[t] = struct{}{}
		miner4.PreCommitSealProofTypesV7[t] = struct{}{}
		miner4.PreCommitSealProofTypesV7[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}
		miner4.PreCommitSealProofTypesV8[t+abi.RegisteredSealProof_StackedDrg2KiBV1_1] = struct{}{}

	}
}

// SetPreCommitChallengeDelay sets the pre-commit challenge delay across all
// actors versions. Use for testing.
func SetPreCommitChallengeDelay(delay abi.ChainEpoch) {
	// Set for all miner versions.

	miner0.PreCommitChallengeDelay = delay

	miner2.PreCommitChallengeDelay = delay

	miner3.PreCommitChallengeDelay = delay

	miner4.PreCommitChallengeDelay = delay

}

// TODO: this function shouldn't really exist. Instead, the API should expose the precommit delay.
func GetPreCommitChallengeDelay() abi.ChainEpoch {
	return miner4.PreCommitChallengeDelay
}

// SetConsensusMinerMinPower sets the minimum power of an individual miner must
// meet for leader election, across all actor versions. This should only be used
// for testing.
func SetConsensusMinerMinPower(p abi.StoragePower) {

	power0.ConsensusMinerMinPower = p

	for _, policy := range builtin2.SealProofPolicies {
		policy.ConsensusMinerMinPower = p
	}

	for _, policy := range builtin3.PoStProofPolicies {
		policy.ConsensusMinerMinPower = p
	}

	for _, policy := range builtin4.PoStProofPolicies {
		policy.ConsensusMinerMinPower = p
	}

}

// SetMinVerifiedDealSize sets the minimum size of a verified deal. This should
// only be used for testing.
func SetMinVerifiedDealSize(size abi.StoragePower) {

	verifreg0.MinVerifiedDealSize = size

	verifreg2.MinVerifiedDealSize = size

	verifreg3.MinVerifiedDealSize = size

	verifreg4.MinVerifiedDealSize = size

}

func GetMaxProveCommitDuration(ver actors.Version, t abi.RegisteredSealProof) abi.ChainEpoch {
	switch ver {

	case actors.Version0:

		return miner0.MaxSealDuration[t]

	case actors.Version2:

		return miner2.MaxProveCommitDuration[t]

	case actors.Version3:

		return miner3.MaxProveCommitDuration[t]

	case actors.Version4:

		return miner4.MaxProveCommitDuration[t]

	default:
		panic("unsupported actors version")
	}
}

func DealProviderCollateralBounds(
	size abi.PaddedPieceSize, verified bool,
	rawBytePower, qaPower, baselinePower abi.StoragePower,
	circulatingFil abi.TokenAmount, nwVer network.Version,
) (min, max abi.TokenAmount) {
	switch actors.VersionForNetwork(nwVer) {

	case actors.Version0:

		return market0.DealProviderCollateralBounds(size, verified, rawBytePower, qaPower, baselinePower, circulatingFil, nwVer)

	case actors.Version2:

		return market2.DealProviderCollateralBounds(size, verified, rawBytePower, qaPower, baselinePower, circulatingFil)

	case actors.Version3:

		return market3.DealProviderCollateralBounds(size, verified, rawBytePower, qaPower, baselinePower, circulatingFil)

	case actors.Version4:

		return market4.DealProviderCollateralBounds(size, verified, rawBytePower, qaPower, baselinePower, circulatingFil)

	default:
		panic("unsupported actors version")
	}
}

func DealDurationBounds(pieceSize abi.PaddedPieceSize) (min, max abi.ChainEpoch) {
	return market4.DealDurationBounds(pieceSize)
}

// Sets the challenge window and scales the proving period to match (such that
// there are always 48 challenge windows in a proving period).
func SetWPoStChallengeWindow(period abi.ChainEpoch) {

	miner0.WPoStChallengeWindow = period
	miner0.WPoStProvingPeriod = period * abi.ChainEpoch(miner0.WPoStPeriodDeadlines)

	miner2.WPoStChallengeWindow = period
	miner2.WPoStProvingPeriod = period * abi.ChainEpoch(miner2.WPoStPeriodDeadlines)

	miner3.WPoStChallengeWindow = period
	miner3.WPoStProvingPeriod = period * abi.ChainEpoch(miner3.WPoStPeriodDeadlines)

	// by default, this is 2x finality which is 30 periods.
	// scale it if we're scaling the challenge period.
	miner3.WPoStDisputeWindow = period * 30

	miner4.WPoStChallengeWindow = period
	miner4.WPoStProvingPeriod = period * abi.ChainEpoch(miner4.WPoStPeriodDeadlines)

	// by default, this is 2x finality which is 30 periods.
	// scale it if we're scaling the challenge period.
	miner4.WPoStDisputeWindow = period * 30

}

func GetWinningPoStSectorSetLookback(nwVer network.Version) abi.ChainEpoch {
	if nwVer <= network.Version3 {
		return 10
	}

	// NOTE: if this ever changes, adjust it in a (*Miner).mineOne() logline as well
	return ChainFinality
}

func GetMaxSectorExpirationExtension() abi.ChainEpoch {
	return miner4.MaxSectorExpirationExtension
}

// TODO: we'll probably need to abstract over this better in the future.
func GetMaxPoStPartitions(p abi.RegisteredPoStProof) (int, error) {
	sectorsPerPart, err := builtin4.PoStProofWindowPoStPartitionSectors(p)
	if err != nil {
		return 0, err
	}
	return int(miner4.AddressedSectorsMax / sectorsPerPart), nil
}

func GetDefaultSectorSize() abi.SectorSize {
	// supported sector sizes are the same across versions.
	szs := make([]abi.SectorSize, 0, len(miner4.PreCommitSealProofTypesV8))
	for spt := range miner4.PreCommitSealProofTypesV8 {
		ss, err := spt.SectorSize()
		if err != nil {
			panic(err)
		}

		szs = append(szs, ss)
	}

	sort.Slice(szs, func(i, j int) bool {
		return szs[i] < szs[j]
	})

	return szs[0]
}

func GetSectorMaxLifetime(proof abi.RegisteredSealProof, nwVer network.Version) abi.ChainEpoch {
	if nwVer <= network.Version10 {
		return builtin4.SealProofPoliciesV0[proof].SectorMaxLifetime
	}

	return builtin4.SealProofPoliciesV11[proof].SectorMaxLifetime
}

func GetAddressedSectorsMax(nwVer network.Version) int {
	switch actors.VersionForNetwork(nwVer) {

	case actors.Version0:
		return miner0.AddressedSectorsMax

	case actors.Version2:
		return miner2.AddressedSectorsMax

	case actors.Version3:
		return miner3.AddressedSectorsMax

	case actors.Version4:
		return miner4.AddressedSectorsMax

	default:
		panic("unsupported network version")
	}
}

func GetDeclarationsMax(nwVer network.Version) int {
	switch actors.VersionForNetwork(nwVer) {

	case actors.Version0:

		// TODO: Should we instead panic here since the concept doesn't exist yet?
		return miner0.AddressedPartitionsMax

	case actors.Version2:

		return miner2.DeclarationsMax

	case actors.Version3:

		return miner3.DeclarationsMax

	case actors.Version4:

		return miner4.DeclarationsMax

	default:
		panic("unsupported network version")
	}
}
