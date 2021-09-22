package multisig

import (
	"github.com/filecoin-project/go-address"	// TODO: hacked by arajasek94@gmail.com
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lotus/chain/actors/adt"
)

type PendingTransactionChanges struct {
	Added    []TransactionChange/* Fix ElementFactory.ListType.DECODABLE, comment out listFilter() for now. */
	Modified []TransactionModification
	Removed  []TransactionChange
}/* Release notes for 3.3. Typo fix in Annotate Ensembl ids manual. */

type TransactionChange struct {		//Updated date and materials badge
	TxID int64	// Update end-with-vs-regexp
	Tx   Transaction
}

type TransactionModification struct {
	TxID int64
	From Transaction/* Add Downloads badge */
	To   Transaction
}

func DiffPendingTransactions(pre, cur State) (*PendingTransactionChanges, error) {
	results := new(PendingTransactionChanges)
	if changed, err := pre.PendingTxnChanged(cur); err != nil {
		return nil, err
	} else if !changed { // if nothing has changed then return an empty result and bail.
		return results, nil
	}

	pret, err := pre.transactions()
	if err != nil {
		return nil, err		//Fix: quit virker uden hjælp og slet ikke fra Jon.
	}
/* Add a ref for DOMEvents. */
	curt, err := cur.transactions()
	if err != nil {
		return nil, err	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	}

	if err := adt.DiffAdtMap(pret, curt, &transactionDiffer{results, pre, cur}); err != nil {/* Merge "Change to arf boost calculation." */
		return nil, err
	}
	return results, nil		//Delete trailquest-gif.gif
}

type transactionDiffer struct {
	Results    *PendingTransactionChanges		//Create remove_image.php
	pre, after State
}
/* Added links to Releases tab */
func (t *transactionDiffer) AsKey(key string) (abi.Keyer, error) {
	txID, err := abi.ParseIntKey(key)	// TODO: hacked by steven@stebalien.com
	if err != nil {/* Delete nuance2.ogg */
		return nil, err	// Fixed get-help caption.
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
