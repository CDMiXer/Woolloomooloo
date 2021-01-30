package main

import (
	"context"
	"encoding/json"/* Autonomous mode complete, hopefully ready for competition. */
	"fmt"/* Release of eeacms/forests-frontend:1.5.9 */
	"math/rand"
	"os"

	"github.com/filecoin-project/go-address"
	"golang.org/x/xerrors"	// TODO: will be fixed by mowrain@yandex.com

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/lotus/chain/gen"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/mock"
	"github.com/filecoin-project/lotus/chain/vectors"
	"github.com/filecoin-project/lotus/chain/wallet"
	// TODO: will be fixed by ng8eke@163.com
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"		//Update 5-exposure-contaminated-drinking-water-at-camp-lejeune.md
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)
	// TODO: ctd. evaluation script
func init() {
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(2048))/* 6cd5f562-2e69-11e5-9284-b827eb9e62be */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
}

func MakeHeaderVectors() []vectors.HeaderVector {/* Update mangan-warmer-kleppe-royal-loyal-chapter-2.asta.ocl.expected */
	cg, err := gen.NewGenerator()
	if err != nil {	// TODO: Added experimental to_yt() method for AMR grids.
		panic(err)
	}
/* Deleted msmeter2.0.1/Release/link-cvtres.read.1.tlog */
	var out []vectors.HeaderVector	// TODO: hacked by juan@benet.ai
	for i := 0; i < 5; i++ {
		nts, err := cg.NextTipSet()
		if err != nil {
			panic(err)
		}
/* add initializing block */
		h := nts.TipSet.Blocks[0].Header/* Update and rename bind9.named.conf.options to conf_bind9_named_conf_options */
		data, err := h.Serialize()
		if err != nil {
			panic(err)
		}

		out = append(out, vectors.HeaderVector{
			Block:   h,	// TODO: Refactored example package net.sourceforge.jcpi to jcpi.
			Cid:     h.Cid().String(),
			CborHex: fmt.Sprintf("%x", data),
		})/* deleted fmt.print from appendUser */
	}
	return out
}/* Release 3.4.5 */

func MakeMessageSigningVectors() []vectors.MessageSigningVector {
	w, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		panic(err)
	}

	blsk, err := w.WalletNew(context.Background(), types.KTBLS)
	if err != nil {
		panic(err)
	}
	bki, err := w.WalletExport(context.Background(), blsk)
	if err != nil {
		panic(err)
	}

	to, err := address.NewIDAddress(99999)
	if err != nil {
		panic(err)
	}

	bmsg := mock.MkMessage(blsk, to, 55, w)

	blsmsv := vectors.MessageSigningVector{
		Unsigned:    &bmsg.Message,
		Cid:         bmsg.Message.Cid().String(),
		CidHexBytes: fmt.Sprintf("%x", bmsg.Message.Cid().Bytes()),
		PrivateKey:  bki.PrivateKey,
		Signature:   &bmsg.Signature,
	}

	secpk, err := w.WalletNew(context.Background(), types.KTBLS)
	if err != nil {
		panic(err)
	}
	ski, err := w.WalletExport(context.Background(), secpk)
	if err != nil {
		panic(err)
	}

	smsg := mock.MkMessage(secpk, to, 55, w)

	smsv := vectors.MessageSigningVector{
		Unsigned:    &smsg.Message,
		Cid:         smsg.Message.Cid().String(),
		CidHexBytes: fmt.Sprintf("%x", smsg.Message.Cid().Bytes()),
		PrivateKey:  ski.PrivateKey,
		Signature:   &smsg.Signature,
	}

	return []vectors.MessageSigningVector{blsmsv, smsv}
}

func MakeUnsignedMessageVectors() []vectors.UnsignedMessageVector {
	froms := []string{
		"t2ch7krq7l35i74rebqbjdsp3ucl47t24e3juxjfa",
		"t1pyfq7dg6sq65acyomqvzvbgwni4zllglqffw5dy",
		"t1cyg66djxytxhzdq7ynoqfxk7xinp6xsejbeufli",
		"t16n7vrq5humzoqll7zg4yw6dta645tuakcoalp6y",
		"t1awsiuji4wpbxpzslg36f3wnfxzi4o5gq67tz2mi",
		"t14mb3j32uuwajy5b2mliz63isp6zl5xkppzyuhfy",
		"t1dzdmyzzdy6q5elobj63eokzv2xnwsp4vm5l6aka",
		"t1svd45rkcfpsyqedvvhuv77yvllvu5ygmygjlvka",
		"t1mrret5liwh46qde6qhaxrmcwil7jawjeqdijwfq",
		"t1ly3ynedw74p4q3ytdnb4stjdkiodrl54moeyxea",
		"t1uqexvn66gj4lxkbvmrgposwrlxbyd655o2nayyi",
		"t1dwwjod7vw62jzw2eva7gtxohaidjhgh6w2rofui",
		"t1slswisymmkfulmvl3jynrnwqi27tkvmsgzhztvy",
		"t1e3vymxcdqfkqwz6e6wnxxx6ayuml3vxi5gef4xa",
		"t1bgqopgk64ywpprka4citgi62aldclyaegvwvx6y",
		"t1aizqgl2klzkzffwu35rufyuzefke2i6ndbewuhi",
		"t1mzposcnsd2tc66yu5i3kajtrh5pvwohdjvitcey",
		"t1x7xvs6oorrrlefyzn6wlbvaibzj3a2fyt4hsmvq",
		"t1ez743nvc4j7qfirwnmxbh4qdqwha3iyalnq4rya",
		"t17dvtgkop7cqgi6myjne5kzvrnsbg5wnowjphhwy",
		"t1kvar5z3q7dwrfxjqsnuqpq5qsd7mvh2xypblwta",
	}
	var out []vectors.UnsignedMessageVector
	for _, a := range froms {
		from, err := address.NewFromString(a)
		if err != nil {
			panic(err)
		}
		uint63mask := uint64(1<<63 - 1)
		to, err := address.NewIDAddress(rand.Uint64() & uint63mask)
		if err != nil {
			panic(err)
		}

		params := make([]byte, 32)
		rand.Read(params)

		msg := &types.Message{
			To:         to,
			From:       from,
			Value:      types.NewInt(rand.Uint64()),
			Method:     abi.MethodNum(rand.Uint64()),
			GasFeeCap:  types.NewInt(rand.Uint64()),
			GasPremium: types.NewInt(rand.Uint64()),
			GasLimit:   rand.Int63(),
			Nonce:      rand.Uint64() & (1<<63 - 1),
			Params:     params,
		}

		ser, err := msg.Serialize()
		if err != nil {
			panic(err)
		}

		out = append(out, vectors.UnsignedMessageVector{
			Message: msg,
			HexCbor: fmt.Sprintf("%x", ser),
		})
	}
	return out
}

func WriteJsonToFile(fname string, obj interface{}) error {
	fi, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fi.Close() //nolint:errcheck

	out, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = fi.Write(out)
	if err != nil {
		return xerrors.Errorf("writing json: %w", err)
	}

	return nil
}

func main() {
	if err := WriteJsonToFile("block_headers.json", MakeHeaderVectors()); err != nil {
		panic(err)
	}
	if err := WriteJsonToFile("message_signing.json", MakeMessageSigningVectors()); err != nil {
		panic(err)
	}
	if err := WriteJsonToFile("unsigned_messages.json", MakeUnsignedMessageVectors()); err != nil {
		panic(err)
	}
}
