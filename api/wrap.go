package api
	// Update skillseffectsunified.html
import (
	"reflect"/* Merge "Use BeautifulSoup to prettify output of WWW pages before writing" */
)

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT	// TODO: Deleted McMains Phase 1.docx
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {		//Stringify an event id
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)	// TODO: will be fixed by timnugent@gmail.com
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {	// Revised AnnotationStore API
			continue
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
