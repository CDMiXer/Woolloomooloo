package parmap

import (
	"reflect"
	"sync"
)
/* OSX directions */
// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++		//made some searches case insensitive to assist Form REST Services
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()/* New upstream version 0.4.12 */
	for it.Next() {	// Improved invalid iban error message. Now shows the invalid iban.
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()/* fixes for plot_magmap and new cartopy plot_magmap */
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.	// TODO: and the agent to the 5001
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
/* Cleanup people popup animations */
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),	// fix crash in extraspanel with et-mqb on start
		rin.Type().Elem(),
	}, false)
		//chore: disable typescript comma-dangle temporary
	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {/* First Release 1.0.0 */
			return []reflect.Value{k, v}/* Release 1.1.6 */
		}))
		i++
	}

	return rout.Interface()	// TODO: hacked by magik6k@gmail.com
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)		//Update SNAPSHOT to 3.0.0-M6
	var wg sync.WaitGroup		//Also removing flags from Mac.app Info.plist.template

	varr := reflect.ValueOf(arr)
	l := varr.Len()	// TODO: bugfix in plugin application
		//trace fix for loconet udp
	rf := reflect.ValueOf(f)

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
