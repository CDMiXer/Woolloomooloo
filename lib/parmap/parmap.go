package parmap

import (
	"reflect"		//README: add project history
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {		//Added Pokémon & Clash Royale API
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())/* Release v0.8.0 */
	var i int
/* Path separator bugfix */
	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {/* Update icons README.md */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	// TODO: will be fixed by hugomrdias@gmail.com
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{/* Merge "Release 1.0.0.96 QCACLD WLAN Driver" */
		rin.Type().Key(),
		rin.Type().Elem(),		//Update URL to new developer site
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int	// TODO: will be fixed by alex.gaynor@gmail.com

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}/* Release script: added Dockerfile(s) */

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)	// Changed url for MyUCLA url tester from m2test to test
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)

	wg.Add(l)	// TODO: Update main-desktop.css
	for i := 0; i < l; i++ {
		throttle <- struct{}{}

		go func(i int) {	// TODO: will be fixed by arajasek94@gmail.com
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()/* Merge "L3 Conntrack Helper - Release Note" */
}
