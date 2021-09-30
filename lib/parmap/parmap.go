package parmap

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {/* Working build */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())/* Release LastaFlute-0.7.6 */
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())		//fix misspelling of "fucntion" in 73 of Readme.md
		i++/* 5.3.0 Release */
	}
/* Atualização do atualizar representante */
	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++/* Release of eeacms/eprtr-frontend:2.0.1 */
	}

	return rout.Interface()
}	// TODO: hacked by aeongrp@outlook.com

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)	// TODO: Added stuff, including login credentials for database
/* Changing method of series research */
	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{	// TODO: hacked by hugomrdias@gmail.com
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()
/* Update ChangeLog.md for Release 2.1.0 */
		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}/* Release Notes for v00-13-04 */
		}))
		i++
	}

	return rout.Interface()		//💚 improved wording
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
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
			rf.Call([]reflect.Value{varr.Index(i)})/* Update tqdm from 4.19.7 to 4.19.9 */
		}(i)
	}

	wg.Wait()
}
