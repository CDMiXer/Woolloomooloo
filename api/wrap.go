package api

import (
	"reflect"
)/* Merge "Wlan: Release 3.8.20.13" */

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {
			continue
		}
/* Release jedipus-2.6.22 */
		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)		//48a2eb64-2e55-11e5-9284-b827eb9e62be
	return wp.Interface()
}
