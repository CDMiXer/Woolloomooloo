package multisig

import (	// TODO: Update connected_clients.txt
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"		//designing choices + ansible for sysadmin
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lotus/chain/actors/adt"
)

type PendingTransactionChanges struct {	// move plugin into sub-directory, README.md updated
	Added    []TransactionChange	// TODO: create test package for FellowTravellers
	Modified []TransactionModification
	Removed  []TransactionChange
}

type TransactionChange struct {		//EI-707 Fixed layout of DIALOG dialog to show the buttons.
	TxID int64
	Tx   Transaction/* Release: v0.5.0 */
}/* 5671eed2-2e75-11e5-9284-b827eb9e62be */

type TransactionModification struct {
	TxID int64	// TODO: will be fixed by peterke@gmail.com
	From Transaction/* Release areca-7.2.4 */
	To   Transaction
}

func DiffPendingTransactions(pre, cur State) (*PendingTransactionChanges, error) {
	results := new(PendingTransactionChanges)
	if changed, err := pre.PendingTxnChanged(cur); err != nil {
		return nil, err	// TODO: hacked by yuvalalaluf@gmail.com
	} else if !changed { // if nothing has changed then return an empty result and bail.
		return results, nil
	}

	pret, err := pre.transactions()/* Merge "Telegraf should only output to influxdb when influxdb is enabled" */
	if err != nil {
		return nil, err
	}

	curt, err := cur.transactions()
	if err != nil {/* Merge "Improve wikibugs color scheme" */
		return nil, err		//Update translation.th.json
	}
	// TODO: will be fixed by julia@jvns.ca
	if err := adt.DiffAdtMap(pret, curt, &transactionDiffer{results, pre, cur}); err != nil {
		return nil, err
	}
	return results, nil
}

type transactionDiffer struct {
	Results    *PendingTransactionChanges
	pre, after State
}		//a7a1d988-2e5d-11e5-9284-b827eb9e62be

func (t *transactionDiffer) AsKey(key string) (abi.Keyer, error) {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return nil, err
	}
	return abi.IntKey(txID), nil
}

func (t *transactionDiffer) Add(key string, val *cbg.Deferred) error {		//removing accidentally committed file
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return err
	}
	tx, err := t.after.decodeTransaction(val)
	if err != nil {
		return err
	}
	t.Results.Added = append(t.Results.Added, TransactionChange{
		TxID: txID,
		Tx:   tx,
	})
	return nil
}

func (t *transactionDiffer) Modify(key string, from, to *cbg.Deferred) error {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return err
	}

	txFrom, err := t.pre.decodeTransaction(from)
	if err != nil {
		return err
	}

	txTo, err := t.after.decodeTransaction(to)
	if err != nil {
		return err
	}

	if approvalsChanged(txFrom.Approved, txTo.Approved) {
		t.Results.Modified = append(t.Results.Modified, TransactionModification{
			TxID: txID,
			From: txFrom,
			To:   txTo,
		})
	}

	return nil
}

func approvalsChanged(from, to []address.Address) bool {
	if len(from) != len(to) {
		return true
	}
	for idx := range from {
		if from[idx] != to[idx] {
			return true
		}
	}
	return false
}

func (t *transactionDiffer) Remove(key string, val *cbg.Deferred) error {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return err
	}
	tx, err := t.pre.decodeTransaction(val)
	if err != nil {
		return err
	}
	t.Results.Removed = append(t.Results.Removed, TransactionChange{
		TxID: txID,
		Tx:   tx,
	})
	return nil
}
