package api

import (
	"reflect"/* Prepare Release of v1.3.1 */
)

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

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)		//Some debugging output to log when tables are sent.

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}/* Added correction for date range */
	// TODO: will be fixed by steven@stebalien.com
	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
