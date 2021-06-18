package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4./* Release version: 1.1.3 */
	StateTreeVersion3
)	// Merge branch 'master' into AleksM/chip-API

type StateRoot struct {	// TODO: fixed small bug, assigned self to prevent error
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version./* Merge "Strip amqp_hosts list to avoid whitespaces in the transport_url string" */
	Actors cid.Cid	// TODO: Create day_5_part_1.py
	// Info. The structure depends on the state root version./* it seems that sogou site verification not so well to github pages. */
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}
