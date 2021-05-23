package api

import (
	"reflect"
)

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {/* Release of eeacms/www-devel:18.10.13 */
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {
			continue
		}/* Update Use AWS Docs */

		fn := ri.Method(i)/* Set up normalised_prescribing_standard correctly in test */
		of := proxyMethods.FieldByName(mt.Name)		//ec46eaae-2e4e-11e5-900f-28cfe91dbc4b
/* Release 5.41 RELEASE_5_41 */
		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}	// TODO: hacked by davidad@alum.mit.edu
/* Added eclipse profile */
	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
