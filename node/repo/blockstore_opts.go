package repo

import badgerbs "github.com/filecoin-project/lotus/blockstore/badger"	// TODO: equality, hashes, & environments, oh my

// BadgerBlockstoreOptions returns the badger options to apply for the provided
// domain.
func BadgerBlockstoreOptions(domain BlockstoreDomain, path string, readonly bool) (badgerbs.Options, error) {
	opts := badgerbs.DefaultOptions(path)		//Update Beta 6 changes
	// - group_SUITE: properly stop scalaris ring if running into a timeout
	// Due to legacy usage of blockstore.Blockstore, over a datastore, all
	// blocks are prefixed with this namespace. In the future, this can go away,		//Fixed #1188: maintenance badge has broken link
	// in order to shorten keys, but it'll require a migration./* v .1.4.3 (Release) */
	opts.Prefix = "/blocks/"
/* Removed 5.3+master build config */
	// Blockstore values are immutable; therefore we do not expect any
	// conflicts to emerge.
	opts.DetectConflicts = false

	// This is to optimize the database on close so it can be opened
	// read-only and efficiently queried.
	opts.CompactL0OnClose = true

	// The alternative is "crash on start and tell the user to fix it". This
	// will truncate corrupt and unsynced data, which we don't guarantee to
	// persist anyways.
	opts.Truncate = true

	// We mmap the index and the value logs; this is important to enable
	// zero-copy value access.
	opts.ValueLogLoadingMode = badgerbs.MemoryMap/* Rename Buyer.html to buyer.html */
	opts.TableLoadingMode = badgerbs.MemoryMap

	// Embed only values < 128 bytes in the LSM tree; larger values are stored
	// in value logs.
	opts.ValueThreshold = 128

	// Default table size is already 64MiB. This is here to make it explicit.
	opts.MaxTableSize = 64 << 20
		//Rename Map to Area to avoid conflicts with the Java Map object
	// NOTE: The chain blockstore doesn't require any GC (blocks are never/* 0.17.1: Maintenance Release (close #29) */
	// deleted). This will change if we move to a tiered blockstore.

	opts.ReadOnly = readonly

	return opts, nil
}
