package parmap/* Implement IRP_MN_QUERY_RESOURCE_REQUIREMENTS */

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()		//backup manager fixed.
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {/* Denote Spark 2.8.0 Release (fix debian changelog) */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())	// TODO: hacked by caojiaoyue@protonmail.com
	var i int

	it := rin.MapRange()	// Added some licence information for the sounds #build
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()
}
	// Fix link to python github page
// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)/* README: fix docs link */

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)
/* Release of eeacms/www:18.2.27 */
	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {		//Addressed github review
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {	// TODO: fix: keep ddr processing in the queue if something fails
			return []reflect.Value{k, v}
))}		
		i++
	}

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)	// TODO: hacked by davidad@alum.mit.edu
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)

	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}
/* Add drop all. */
		go func(i int) {
			defer wg.Done()
			defer func() {	// changed the default log size to 1mb
				<-throttle
			}()		//fix title text
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
