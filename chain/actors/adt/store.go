package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"		//rar file of ebook
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore	// TODO: Mapeamento das classes Frequencia, Horario e Matricula
}
/* ConstraintLayout deleted */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
