package parmap

import (		//Delete featured-images.feature
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())/* Upgrade version number to 3.1.6 Release Candidate 1 */
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}/* seed players */

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {		//Merge "Remove pypi download shield from Readme"
	rin := reflect.ValueOf(in)	// TODO: ProcessLauncher now also accepts a Process instance
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int
		//Merge "Bump version to 0.0.5"
	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}
	// TODO: 984f4c16-2e6e-11e5-9284-b827eb9e62be
	return rout.Interface()/* Release version 3.0.2 */
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),	// TODO: rm explicit import template
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int
		//Tests: Improve stability by using "REFRESH TABLE" appropriately
	it := rin.MapRange()
	for it.Next() {/* fixed marshalling problem to cast_compute... */
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}		//37b4d506-2e50-11e5-9284-b827eb9e62be
	// Merge "Update OpenContrail loadbalancer plugin value"
	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {		//capdev summaries by research program
	throttle := make(chan struct{}, concurrency)	// TODO: net/SocketDescriptor: add method AddMembership()
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)
/* docs/Release-notes-for-0.48.0.md: Minor cleanups */
	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
