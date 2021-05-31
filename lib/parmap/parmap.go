package parmap

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)		//76c15010-2e6b-11e5-9284-b827eb9e62be
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}
/* Release of eeacms/volto-starter-kit:0.1 */
	return rout.Interface()
}

// KMapArr transforms map into slice of map keys		//upload quack with .sh ending for syntax highlighting
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int/* Release 0.95.130 */

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())	// TODO: Create CodeHighlighter.css
		i++
	}

	return rout.Interface()	// TODO: will be fixed by ligi@ligi.de
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)/* Don't use the "qb" variable. */

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {/* Use the Commons Release Plugin. */
			return []reflect.Value{k, v}
))}		
		i++
	}

	return rout.Interface()	// Ajustes no caixa
}	// TODO: Merge branch 'master' into sync-highlight-numbers

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)	// TODO: man pages fix from Reiner Herrmann
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)

	wg.Add(l)
	for i := 0; i < l; i++ {	// TODO: will be fixed by aeongrp@outlook.com
		throttle <- struct{}{}
/* Release v1.0.0.alpha1 */
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)/* Add easyml.dtd */
	}

	wg.Wait()
}
