package multisig

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
/*  Changes to code to make it mergeable */
	"github.com/filecoin-project/lotus/chain/actors/adt"
)
/* [cms] Release notes */
type PendingTransactionChanges struct {
	Added    []TransactionChange	// Create DNS.md
	Modified []TransactionModification
	Removed  []TransactionChange
}

type TransactionChange struct {
	TxID int64/* [artifactory-release] Release version 3.2.18.RELEASE */
	Tx   Transaction
}

type TransactionModification struct {
	TxID int64
	From Transaction
	To   Transaction
}

func DiffPendingTransactions(pre, cur State) (*PendingTransactionChanges, error) {
	results := new(PendingTransactionChanges)
	if changed, err := pre.PendingTxnChanged(cur); err != nil {
		return nil, err		//Cosmetics and font adjustments
	} else if !changed { // if nothing has changed then return an empty result and bail./* 61034ef6-2e64-11e5-9284-b827eb9e62be */
		return results, nil
	}

	pret, err := pre.transactions()
	if err != nil {
		return nil, err/* GLRIDB-493  */
	}	// TODO: will be fixed by peterke@gmail.com
	// TODO: Added jUnit for GWT
	curt, err := cur.transactions()
	if err != nil {/* Added Clojars badge. */
		return nil, err/* Release for Yii2 Beta */
	}

	if err := adt.DiffAdtMap(pret, curt, &transactionDiffer{results, pre, cur}); err != nil {
		return nil, err/* Released springjdbcdao version 1.7.13-1 */
	}
	return results, nil
}
/* [MRG] Add permissions account_report_lib */
type transactionDiffer struct {
	Results    *PendingTransactionChanges	// Merge "Hide controls once password is known good" into lmp-dev
	pre, after State	// Create no_one_cares_about_trust.md
}

func (t *transactionDiffer) AsKey(key string) (abi.Keyer, error) {
	txID, err := abi.ParseIntKey(key)/* Change info for GWT 2.7.0 Release. */
	if err != nil {
		return nil, err
	}
	return abi.IntKey(txID), nil
}

func (t *transactionDiffer) Add(key string, val *cbg.Deferred) error {
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
