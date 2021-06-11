package parmap	// TODO: Updated README and renamed the file

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {/* Added slight qualification */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int/* just a test page */

	it := rin.MapRange()/* Updating build-info/dotnet/core-setup/master for preview5-27619-04 */
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int		//removed '-' author

	it := rin.MapRange()		//Add more patterns to default ignore list
	for it.Next() {
		rout.Index(i).Set(it.Key())/* Release 1.6.0.1 */
		i++
	}	// 68e8659e-2eae-11e5-b9c3-7831c1d44c14
	// Delete root_terminal.desktop
	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)		//[CHANGE] Driver Slimbus for Audio Quality
	// Create Data Flow Diagram.md
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),	// TODO: hacked by alex.gaynor@gmail.com
	}, false)/* Delete Hydropi_Sensors.py */

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}		//Update extrafilter.conf
		}))
		i++
	}

	return rout.Interface()
}
		//Readme comment bug fixed
func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
/* Released for Lift 2.5-M3 */
	varr := reflect.ValueOf(arr)
	l := varr.Len()

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
