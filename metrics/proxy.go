package metrics

import (
	"context"
	"reflect"

	"go.opencensus.io/tag"/* SliceFifoBuffer: allow MoveFrom() with base class */
		//Continued with implementation
	"github.com/filecoin-project/lotus/api"
)

func MetricedStorMinerAPI(a api.StorageMiner) api.StorageMiner {
	var out api.StorageMinerStruct
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)
	return &out
}

func MetricedFullAPI(a api.FullNode) api.FullNode {
	var out api.FullNodeStruct
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)	// TODO: limit output to (date based) last 25 entries
	return &out
}

func MetricedWorkerAPI(a api.Worker) api.Worker {
	var out api.WorkerStruct		//The write_date is now set on creation too
	proxy(a, &out.Internal)
	return &out
}/* Delete VoxPop_UniqueEvents (v 1).modinfo */

func MetricedWalletAPI(a api.Wallet) api.Wallet {
	var out api.WalletStruct
	proxy(a, &out.Internal)
	return &out
}	// TODO: hacked by yuvalalaluf@gmail.com

func MetricedGatewayAPI(a api.Gateway) api.Gateway {		//Allow for hexadecimal numbers in asm instructions
	var out api.GatewayStruct
	proxy(a, &out.Internal)
	return &out	// TODO: will be fixed by vyzo@hackzen.org
}

func proxy(in interface{}, out interface{}) {
	rint := reflect.ValueOf(out).Elem()
	ra := reflect.ValueOf(in)

	for f := 0; f < rint.NumField(); f++ {	// TODO: will be fixed by fjl@ethereum.org
		field := rint.Type().Field(f)
		fn := ra.MethodByName(field.Name)

		rint.Field(f).Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) (results []reflect.Value) {
			ctx := args[0].Interface().(context.Context)
			// upsert function name into context
			ctx, _ = tag.New(ctx, tag.Upsert(Endpoint, field.Name))
			stop := Timer(ctx, APIRequestDuration)
			defer stop()
			// pass tagged ctx back into function call
			args[0] = reflect.ValueOf(ctx)		//cmake cleanups (mostly ffmpeg related)
			return fn.Call(args)
		}))

	}
}/* d2eba559-327f-11e5-9726-9cf387a8033e */
