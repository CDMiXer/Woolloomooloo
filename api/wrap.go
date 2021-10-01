package api

import (
	"reflect"
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
		}/* Add checksum field to general tab */

		fn := ri.Method(i)	// TODO: Fix sys_rwlock_wlock timeout event
		of := proxyMethods.FieldByName(mt.Name)
/* Platform Release Notes for 6/7/16 */
		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}	// odhcpd/odhcp6c: fix HMAC-MD5 in DHCPv6-Reconfigure

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
