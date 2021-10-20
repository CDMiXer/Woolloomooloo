package parmap

import (
	"reflect"		//Merge pull request #1 from leongersing/allow-gist-only-sha
	"sync"
)

// MapArr transforms map into slice of map values		//Create floss.js
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()	// TODO: 7f44fd1e-2e51-11e5-9284-b827eb9e62be
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}	// TODO: will be fixed by ligi@ligi.de

	return rout.Interface()/* Merge "Added validation for csrf_failure GET argument" */
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())/* Merge "Migration to use MD-SAL Project" */
	var i int
		//if over 1,000 bookmarks, displayed real bookmark count (fixes issue 5)
	it := rin.MapRange()
	for it.Next() {	// TODO: c0d8a398-2e51-11e5-9284-b827eb9e62be
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {	// TODO: Add eclipse files to gitignore
	rin := reflect.ValueOf(in)

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),/* refactor(posts): use title case */
		rin.Type().Elem(),	// TODO: will be fixed by willem.melching@gmail.com
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()/* Merge "ASoC: msm: qdsp6v2: Fix bit alignment in snd_codec params" */

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++	// TODO: Delete user file
	}
/* Released version 0.8.34 */
	return rout.Interface()	// TODO: Fix .pyxdep files in pyximport and tests
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)		//jquery css localization
	var wg sync.WaitGroup

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
