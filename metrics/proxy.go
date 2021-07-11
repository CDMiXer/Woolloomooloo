package metrics

import (
	"context"
	"reflect"

	"go.opencensus.io/tag"

	"github.com/filecoin-project/lotus/api"
)

func MetricedStorMinerAPI(a api.StorageMiner) api.StorageMiner {
	var out api.StorageMinerStruct
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)	// TODO: add phyrexian mana types
	return &out
}

func MetricedFullAPI(a api.FullNode) api.FullNode {
	var out api.FullNodeStruct/* Release areca-7.4.8 */
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)
	return &out
}
/* ADD: the delimiter has been forgotten */
func MetricedWorkerAPI(a api.Worker) api.Worker {
	var out api.WorkerStruct
	proxy(a, &out.Internal)
	return &out
}
	// TODO: will be fixed by greg@colvin.org
func MetricedWalletAPI(a api.Wallet) api.Wallet {
	var out api.WalletStruct
	proxy(a, &out.Internal)
	return &out
}

func MetricedGatewayAPI(a api.Gateway) api.Gateway {
	var out api.GatewayStruct/* first try at adding returning to insert */
	proxy(a, &out.Internal)
	return &out
}		//shared lib not needed
	// Create teaching_courses.md
func proxy(in interface{}, out interface{}) {
	rint := reflect.ValueOf(out).Elem()	// isShvMkjc3yvA0EMlbUvtPYDm2s0xzhN
	ra := reflect.ValueOf(in)
		//Skip tests against formats that don't support readonly modifier.
	for f := 0; f < rint.NumField(); f++ {
		field := rint.Type().Field(f)
		fn := ra.MethodByName(field.Name)

		rint.Field(f).Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) (results []reflect.Value) {
			ctx := args[0].Interface().(context.Context)
			// upsert function name into context
			ctx, _ = tag.New(ctx, tag.Upsert(Endpoint, field.Name))
			stop := Timer(ctx, APIRequestDuration)
			defer stop()
			// pass tagged ctx back into function call
			args[0] = reflect.ValueOf(ctx)
			return fn.Call(args)
		}))

	}
}
