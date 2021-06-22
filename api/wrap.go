package api/* Ensure rake is called within the correct bundle context */

import (
	"reflect"	// TODO: Route-manager distance helpers for Hyde and others working on VNAV support.
)	// TODO: d768ac82-2e62-11e5-9284-b827eb9e62be

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT		//e280cd28-2ead-11e5-bef1-7831c1d44c14
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)		//Dokumentation ergänzt
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {
			continue
		}	// TODO: Merge "Hard refresh the main page on an edit."

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}	// TODO: hacked by fjl@ethereum.org

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
