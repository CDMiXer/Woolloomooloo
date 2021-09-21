package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version./* Update README with coding table examples */
type StateTreeVersion uint64
	// Add implementation of SearchRow.object
const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota	// TODO: hacked by aeongrp@outlook.com
	// StateTreeVersion1 corresponds to actors v2/* Deleted GithubReleaseUploader.dll, GithubReleaseUploader.pdb files */
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2	// TODO: 2TfovIaqRu4NXZ3tIrAA69h7VXroUdJZ
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3	// Merge "Removing vagrant support"
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

.siht noisrev :ODOT //
type StateInfo0 struct{}
